(function(){"use strict";self.importScripts("/gcsim/wasm_exec.js"),WebAssembly.instantiateStreaming||(WebAssembly.instantiateStreaming=async(e,t)=>{const r=await(await e).arrayBuffer();return await WebAssembly.instantiate(r,t)});function n(e){const t=new Go;WebAssembly.instantiateStreaming(fetch(e.wasm),t.importObject).then(r=>{t.run(r.instance),postMessage({type:"ready"})}).catch(r=>{console.error(r),postMessage({type:"failed",reason:r instanceof Error?r.message:"Unknown Error"})})}function s(e){const t=initializeWorker(e.cfg);return t!=null?{type:"failed",reason:JSON.parse(t).error}:{type:"initialized"}}function a(e){try{const t=simulate();return typeof t=="string"||t instanceof String?{type:"failed",reason:JSON.parse(t).error}:{type:"done",result:t,itr:e.itr}}catch(t){return console.log("simulate() call failed"),{type:"failed",reason:`Failed with error: ${t}`}}}function i(e){switch(e.type){case"ready":return n(e);case"initialize":return postMessage(s(e));case"run":return postMessage(a(e));default:throw console.error("aggregator - unknown request: ",e),new Error("aggregator unknown request")}}onmessage=e=>i(e.data)})();
