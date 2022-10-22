import Model3D from "../Model3D";
import React, {useState} from "react";
import PopupModal from "../Modal/PopupModal";
import FormInput from "../FormInput/FormInput";

const Model3DTable = ({model3ds}) => {

  const [model3dInfoModal, setModel3DInfoModal] = useState(false)
//   console.log(model3ds);
    return (
      <>
        <div className="sections-list">
          {model3ds.length && (
              model3ds.map((model3d) => (
                <Model3D showModel3DInfoModal={() => setModel3DInfoModal(model3d)} key={model3d.id} model3d={model3d}/>
              ))
          )}
          {!model3ds.length && (
              <p>No model3ds found!</p>
          )}
        </div>
        {model3dInfoModal && <PopupModal
						modalTitle={"Model3D Info"}
						onCloseBtnPress={() => {
							setModel3DInfoModal(false);
						}}
					>
						<div className="mt-4 text-left">
							<form className="mt-5">
								<FormInput
									disabled
									type={"text"}
									name={"label"}
									label={"Label"}
									value={model3dInfoModal?.label}
								/>
								<FormInput
									disabled
									type={"text"}
									name={"url"}
									label={"Url"}
									value={model3dInfoModal?.url}
								/>
								<FormInput
									disabled
									type={"text"}
									name={"source"}
									label={"Source"}
									value={model3dInfoModal?.source}
								/>
							</form>
						</div>
					</PopupModal>}
      </>
    )
}

export default Model3DTable;