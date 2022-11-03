
import React, {} from "react";

const Model3D = ({ model3d, showModel3DInfoModal}) => {


	var re = /(?:\.([^.]+))?$/;
	var ext = re.exec(model3d.file_name_and_extension)[1];
	var name = model3d.file_name_and_extension.slice(0, model3d.file_name_and_extension.length - ext.length-1);
	console.log(model3d.file_name_and_extension);
	console.log("name: " + name);
	console.log("ext: " + ext);

	function base64ToArrayBuffer(base64) {
		var binaryString = window.atob(base64);
		console.log(binaryString);
		var binaryLen = binaryString.length;
		var bytes = new Uint8Array(binaryLen);
		for (var i = 0; i < binaryLen; i++) {
			var ascii = binaryString.charCodeAt(i);
			bytes[i] = ascii;
		}
		return bytes;
	}

	var sampleArr = base64ToArrayBuffer(model3d.serialized_file_3d);

	let link = document.createElement('a');
	link.download = model3d.file_name_and_extension;
	let blob = new Blob([sampleArr], {type: `model/${ext}`});
	console.log("blob: " + blob);
	link.href = URL.createObjectURL(blob);
	console.log("link.href:" + link.href);

	const RenderObject = () => {

		return <primitive object={model3d.previewFile} scale={1.5}/>
		// URL.revokeObjectURL(link.href);//USE PROMISE?
	};

	return (
		model3d && (
			<>				
				<div
					onLoad={(e) => { showModel3DInfoModal(); e.stopPropagation() }}
					className = "bg-black m-5 w-200 min-h-300 h-3 mb-2 float-left"
				>
					<div className = "bg-white relative inset-y-0 left-0">
						<div className = "text-black ">
							<div>
								<RenderObject/>
							</div>
							<div className="relative">Title: {model3d.title}</div>
							{/* <div className="relative">Description: {model3d.description}</div> */}
							<div className="">Price: {model3d.price}</div>
							{/* <div className="relative">SalePrice: {model3d.price}</div> */}
							<div className="relative">FileName: {model3d.file_name_and_extension}</div>
							<div href={`www.wikipedia.org`}
									onClick={(e) => e.stopPropagation()} > More Details</div>
						</div>
					</div>
				</div>
					{/* <div className="w-full xl:w-1/4 md:w-1/4">
						<h2 className="mb-4 font-semibold tracking-widest text-black uppercase title-font">
						{model3d.title}
							</h2>
						<div className="relative flex flex-col h-full p-8 ">
							<h2 className="mb-4 font-semibold tracking-widest text-black uppercase title-font">
								{model3d?.label}
							
							</h2>
							<h2 className="items-center mb-2 text-lg font-normal tracking-wide text-white">
								<span className="inline-flex items-center justify-center flex-shrink-0 w-5 h-5 mr-2 text-white rounded-full bg-blue-1300">
									<svg
										fill="none"
										stroke="currentColor"
										strokeLinecap="round"
										strokeLinejoin="round"
										strokeWidth="2.5"
										className="w-4 h-4"
										viewBox="0 0 24 24"
									>
										<path d="M20 6L9 17l-5-5"></path>
									</svg>
								</span>
								{model3d?.source}
							</h2>
						</div>
					</div>
					<div className="w-full xl:w-1/4 md:w-1/2 lg:ml-auto" style={{ zindex: 10000 }}>
						<div className="relative flex flex-col h-full p-8">
							<h1 className="flex items-end mx-auto text-3xl font-black leading-none text-white ">
								<span>View Model3D </span>
							</h1>
							<div className="flex flex-col md:flex-row">

								<a
									href={`${model3d?.url}`}
									onClick={(e) => e.stopPropagation()}
									className="transition ease-in-out delay-150 hover:-translate-y-1 hover:scale-110  bg-teal-600 cursor-pointer hover:bg-teal-700 text-white font-bold px-4 py-2 mx-auto mt-3 rounded"
								>
								</a>
							</div>
						</div>
					</div>
				</div> */}
			</>
		)
	);
};

export default Model3D;
