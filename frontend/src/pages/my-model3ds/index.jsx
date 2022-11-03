import React, { useEffect, useState } from "react";
import FiberClient from "../../client";
import config from "../../config";
import DashboardHeader from "../../components/DashboardHeader";
import Footer from "../../components/Footer";
import jwtDecode from "jwt-decode";
import * as moment from "moment";
import Model3DTable from "../../components/Model3DTable";
import FormInput from "../../components/FormInput/FormInput";
import Button from "../../components/Button/Button";
import { NotLoggedIn } from "../not-logged-in/NotLoggedIn";
import Loader from "../../components/Loader";
import PopupModal from "../../components/Modal/PopupModal";
import Buffer from "Buffer";

const client = new FiberClient(config);

const ProfileView = ({ model3ds }) => {
	return (
		<>
			<Model3DTable
				model3ds={model3ds}
				
				showUpdate={true}
			/>
			
		</>
	);
};

const Model3DDashboard = () => {
	const [isLoggedIn, setIsLoggedIn] = useState(false);
	const [error, setError] = useState({ label: "", url: "", source: "" });
	// const [model3DForm, setModel3DForm] = useState({
	// 	label: "",
	// 	url: "https://",
	// 	source: "",
	// });

	const [model3DForm, setModel3DForm] = useState({
        title: "",
        description: "",
        price: "",
        blobData: null,
        fileName: ""
    });

	const [showForm, setShowForm] = useState(false);
	const [model3ds, setModel3Ds] = useState([]);

	function addItem(item) { 
		setModel3Ds((model3ds) => [...model3ds, item])
	}

	const [loading, setLoading] = useState(false);
	const [refreshing, setRefreshing] = useState(true);

	useEffect(() => {
		fetchUserModel3Ds();
	}, []);

	const fetchUserModel3Ds = () => {
		client.getUserModel3Ds().then((data) => {
			console.log("client.getUserModel3Ds -> data.length: " + data.length);
			setRefreshing(false);
			data.forEach(model3d => {
				model3d.blobData = Buffer.from(model3d.serialized_file_3d);
				addItem(model3d);
			});
			
			// setModel3Ds(data?.results);
		});
		console.log("model3ds.length: " + model3ds.length);
	};

   const urlPatternValidation = URL => {
        const regex = new RegExp('(https?://)?([\\da-z.-]+)\\.([a-z.]{2,6})[/\\w .-]*/?');
          return regex.test(URL);
        };

	const onCreateModel3D = (e) => {
		e.preventDefault();
		setLoading(true);
		setError(false);

		if (model3DForm.label.length <= 0) {
			setLoading(false);
			return setError({ label: "Please Enter Model3D Label" });
		}
		if (model3DForm.url.length <= 0) {
			setLoading(false);
			return setError({ url: "Please Enter Model3D Url" });
		}
		if (!urlPatternValidation(model3DForm.url)) {
			setLoading(false);
			return setError({ url: "Please Enter Valid URL" });
		}
		if (model3DForm.source.length <= 0) {
			setLoading(false);
			return setError({ source: "Please Enter Model3D Source" });
		}

		client.fetchUser().then((user) => {
			client
				.createModel3D(
					model3DForm.label,
					model3DForm.url,
					model3DForm.source,
					user?.id
				)
				// eslint-disable-next-line no-unused-vars
				.then((data) => {  // eslint:ignore
					fetchUserModel3Ds();
					setLoading(false);
					setShowForm(false);
				});
		});
	};

	useEffect(() => {
		const tokenString = localStorage.getItem("token");
		if (tokenString) {
			const decodedAccessToken = jwtDecode(tokenString);
			if(moment.unix(decodedAccessToken.exp).toDate() > new Date()) {
				setIsLoggedIn(true);
			}
		}
	}, [])


	if (refreshing) return !isLoggedIn ? <NotLoggedIn /> : <Loader />;

	return (
		<>
			<section
				className="flex flex-col bg-black text-center"
				style={{ minHeight: "100vh" }}
			>
				<DashboardHeader />
				<div className="container px-5 pt-6 text-center mx-auto lg:px-20">
						{/*TODO - move to component*/}
					<h1 className="mb-12 text-3xl font-medium text-white">
						Model3Ds - Better than all the REST
					</h1>

					<button
						className="my-5 text-white bg-teal-500 p-3 rounded"
						onClick={() => {
							setShowForm(!showForm);
						}}
					>
						Create Model3D
					</button>

					<p className="text-base leading-relaxed text-white">Latest model3ds</p>
					<div className="mainViewport text-white">
						{model3ds.length && (
							<ProfileView
								model3ds={model3ds}
								fetchUserModel3Ds={fetchUserModel3Ds}
							/>
						)}
					</div>
				</div>

				<Footer />
			</section>
			{showForm && (
				<PopupModal
					modalTitle={"Create Model3D"}
					onCloseBtnPress={() => {
						setShowForm(false);
						setError({ fullName: "", email: "", password: "" });
					}}
				>
					<div className="mt-4 text-left">
						<form className="mt-5" onSubmit={(e) => onCreateModel3D(e)}>
							<FormInput
								type={"text"}
								name={"label"}
								label={"Label"}
								error={error.label}
								value={model3DForm.label}
								onChange={(e) =>
									setModel3DForm({ ...model3DForm, label: e.target.value })
								}
							/>
							<FormInput
								type={"text"}
								name={"url"}
								label={"Url"}
								error={error.url}
								value={model3DForm.url}
								onChange={(e) =>
									setModel3DForm({ ...model3DForm, url: e.target.value })
								}
							/>
							<FormInput
								type={"text"}
								name={"source"}
								label={"Source"}
								error={error.source}
								value={model3DForm.source}
								onChange={(e) =>
									setModel3DForm({ ...model3DForm, source: e.target.value })
								}
							/>
							<Button
								loading={loading}
								error={error.source}
								title={"Create Model3D"}
							/>
						</form>
					</div>
				</PopupModal>
			)}
		</>
	);
};

export default Model3DDashboard;
