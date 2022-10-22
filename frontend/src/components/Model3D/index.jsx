
import React, { Suspense } from "react";
// import * as THREE from 'three';

import { Canvas, useLoader } from '@react-three/fiber'
// import { Environment, OrbitControls } from "@react-three/drei";

// import { DDSLoader } from "three-stdlib";

// import {OrbitControls} from 'https://threejsfundamentals.org/threejs/resources/threejs/r127/examples/jsm/controls/OrbitControls.js';
import {OBJLoader} from 'https://threejsfundamentals.org/threejs/resources/threejs/r127/examples/jsm/loaders/OBJLoader.js';
// import { OBJLoader2 } from 'https://threejsfundamentals.org/threejs/resources/threejs/r115/examples/jsm/loaders/OBJLoader2.js';
// import {MTLLoader} from 'https://threejsfundamentals.org/threejs/resources/threejs/r127/examples/jsm/loaders/MTLLoader.js';
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader';

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
	// const loader = new OBJLoader2();
	const RenderObject = () => {
		

		// loader.load(link.href, function ( obj ) {
		// 	console.log(obj);
		// 	// setModelFile(obj);
		// }, undefined, function ( error ) {
		// 	console.error( error );
		// } );
		
		const { scene } = ext == "glb" ? useLoader(GLTFLoader, link.href) : useLoader(OBJLoader, link.href);
		return <primitive object={scene} scale={1.5}/>
		// URL.revokeObjectURL(link.href);//USE PROMISE?
	};
	//https://stackoverflow.com/questions/35038884/download-obj-from-bytes-in-javascript
//https://mockstacks.com/How-to-Convert-Blob-to-File-Using-Javascript

	return (
		model3d && (
			<>				
				<div
					onLoad={(e) => { showModel3DInfoModal(); e.stopPropagation() }}
					className="flex flex-wrap items-end justify-between w-full transition duration-500 ease-in-out transform bg-white border-2 border-gray-600 rounded-lg hover:border-white mb-3"
				>
					<div className = "relative inset-y-0 left-0">
						<div className = "text-black">
							<h1>Title: {model3d.title}</h1>
							<h1>Description: {model3d.description}</h1>
							<h1>Price: {model3d.price}</h1>
							<h1>FileName: {model3d.file_name_and_extension}</h1>
						</div>
						<Canvas>
							<ambientLight />
							<Suspense fallback={null}>
								<RenderObject />
							</Suspense>
						</Canvas>
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
