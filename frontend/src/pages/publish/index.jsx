import React, { useEffect, useState } from "react";
import FiberClient from "../../client";
import config from "../../config";
import DashboardHeader from "../../components/DashboardHeader";
// import Footer from "../../components/Footer";
import jwtDecode from "jwt-decode";
import * as moment from "moment";
// import Model3DTable from "../../components/Model3DTable";
import FormInput from "../../components/FormInput/FormInput";
import Button from "../../components/Button/Button";
// import { NotLoggedIn } from "../not-logged-in/NotLoggedIn";
// import Loader from "../../components/Loader";
// import PopupModal from "../../components/Modal/PopupModal";


const client = new FiberClient(config);

const Publish = () => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [error, setError] = useState({ label: "", url: "", source: "" });
    const [loading, setLoading] = useState(false);
    // const [uploadFile, setUploadFile] = useState();
    // const [refreshing, setRefreshing] = useState(false);
    // setRefreshing(false);
    
    const [model3DForm, setModel3DForm] = useState({
        title: "",
        description: "",
        price: "",
        previewBlobData: null,
        previewFileName: "",
        blobData: null,
        fileName: ""
    });

    function ByteArrayFromFileLoadEvent(event) {
      const reader = new FileReader();
      const fileByteArray = [];
      reader.readAsArrayBuffer(event.target.files[0]);
      reader.onloadend = function (evt) {
          if (evt.target.readyState == FileReader.DONE) {
             const arrayBuffer = evt.target.result,
                 array = new Uint8Array(arrayBuffer);
             for (var i = 0; i < array.length; i++) {
                 fileByteArray.push(array[i]);
              }
          }
      }
      return fileByteArray;
    }

    function onPreviewFileChange(event) {
      const fileByteArray = ByteArrayFromFileLoadEvent(event);
      
      setModel3DForm(previousState => {
        return {...previousState, previewFileName : event.target.files[0].name, previewBlobData : fileByteArray}
      });
    }


    function onFileChange(event) {
      const fileByteArray = ByteArrayFromFileLoadEvent(event);
      
      setModel3DForm(previousState => {
        return {...previousState, fileName : event.target.files[0].name, blobData : fileByteArray}
      });
    }
  
    
    const onPublish = (e) => {
      console.log("onPublish");
      e.preventDefault();
      setError(false);
      setLoading(true);
      console.log(localStorage.getItem('token'))
      console.log(parseFloat(model3DForm.price).toFixed(2))

      client.fetchUser().then((user) => {
        // console.log("user[0].id: " + user[0].id)
        // console.log("model3DForm.fileName: " + model3DForm.fileName)
              client
                  .createModel3D(
            model3DForm.title,
            user.id,
            model3DForm.description,
            model3DForm.price,
            model3DForm.previewBlobData,
            model3DForm.previewFileName,
            model3DForm.blobData,
            model3DForm.fileName,
                  )
                  // eslint-disable-next-line no-unused-vars
                  .then((data) => {  // eslint:ignore
                      setLoading(false);
                  });
          });

      console.log(model3DForm);
      if (model3DForm.title.length <= 0) {
              setLoading(false);
              return setError({ title: "Please Enter Publishing Title" });
      }
      if (model3DForm.description.length <= 0) {
              setLoading(false);
              return setError({ description: "Please Enter Publishing Description" });
      }
      if (model3DForm.price.length <= 0) {
              setLoading(false);
              return setError({ price: "Please Enter Publishing Price" });
      }
      if (model3DForm.blobData.length <= 0) {
              setLoading(false);
              return setError({ blobData: "Please Enter Publishing file" });
      }
      if (model3DForm.fileName.length <= 0) {//todo check for proper extension
              setLoading(false);
              return setError({ fileName: "Please Enter Publishing File Name" });
      }
    };

    useEffect(() => {
      const tokenString = localStorage.getItem("token");
      console.log("y u no execute");
      if (tokenString) {
        const decodedAccessToken = jwtDecode(tokenString);
        console.log("token date: " + moment.unix(decodedAccessToken.exp).toDate());
        console.log("new date: " + (new Date()));
        if (moment.unix(decodedAccessToken.exp).toDate() > new Date()) {
           setIsLoggedIn(true);
        }
        if(isLoggedIn) {
          console.log("user is logged in!");
        }
        else {
          console.log("WARNING:user is not logged in");
        }
      }
    }, []);
    console.log("localStorage.getItem(\"token\"): " + localStorage.getItem("token"));
    
    // if (refreshing) return !isLoggedIn ? <NotLoggedIn /> : <Loader />;
    
    return (
      <>
      <DashboardHeader/>
        <div className="bg-white">
          <h3>
            Publishing!
          </h3>
          <div className="mt-4 text-left">
						<form className="mt-5" onSubmit={(e) => onPublish(e)}>
              
							<FormInput
								type={"text"}
								name={"title"}
								label={"Title"}
								error={error.title}
								value={model3DForm.title}
								onChange={(e) => {
                  setModel3DForm(previousState => {
                    return {...previousState, title : e.target.value}
                  });
								}}
							/>
							<FormInput
								type={"text"}
								name={"description"}
								label={"Description"}
								error={error.description}
								value={model3DForm.description}
								onChange={(e) => {
                  setModel3DForm(previousState => {
                    return {...previousState, description : e.target.value}
                  });
								}}
							/>
              <label htmlFor="previewFile">Preview Image File: </label>
              <input type="file" className="w-72 shadow rounded" id="previewFile"
                onChange={onPreviewFileChange} 
                />
              <label htmlFor="model3dFile">3D Model File: </label>
              <input type="file" className="w-72 shadow rounded" id="model3dFile"
                onChange={onFileChange} 
                />
              <FormInput
								type={"text"}
								name={"price"}
								label={"Price"}
								error={error.price}
								value={model3DForm.price}
								onChange={(e) => {
                  setModel3DForm(previousState => {
                    return {...previousState, price : e.target.value}
                  });
								}}
							/>
							<Button
								loading={loading}
								error={error.source}
								title={"Publish"}
                onClick={onPublish}
							/>
						</form>
					</div>
          <FormInput  
            error={error.source}>
            
            </FormInput>
      </div>
      </>
    );
};

export default Publish;