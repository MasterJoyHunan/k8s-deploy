import{r as d,o as v}from"./index-45490d46.js";/*! *****************************************************************************
Copyright (c) Microsoft Corporation. All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at http://www.apache.org/licenses/LICENSE-2.0

THIS CODE IS PROVIDED ON AN *AS IS* BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, EITHER EXPRESS OR IMPLIED, INCLUDING WITHOUT LIMITATION ANY IMPLIED
WARRANTIES OR CONDITIONS OF TITLE, FITNESS FOR A PARTICULAR PURPOSE,
MERCHANTABLITY OR NON-INFRINGEMENT.

See the Apache Version 2.0 License for specific language governing permissions
and limitations under the License.
***************************************************************************** */var p=function(t,e){return p=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(n,o){n.__proto__=o}||function(n,o){for(var r in o)o.hasOwnProperty(r)&&(n[r]=o[r])},p(t,e)};function y(t,e){p(t,e);function n(){this.constructor=t}t.prototype=e===null?Object.create(e):(n.prototype=e.prototype,new n)}function w(t){var e=typeof Symbol=="function"&&t[Symbol.iterator],n=0;return e?e.call(t):{next:function(){return t&&n>=t.length&&(t=void 0),{value:t&&t[n++],done:!t}}}}function E(t,e){var n=typeof Symbol=="function"&&t[Symbol.iterator];if(!n)return t;var o=n.call(t),r,i=[],s;try{for(;(e===void 0||e-- >0)&&!(r=o.next()).done;)i.push(r.value)}catch(c){s={error:c}}finally{try{r&&!r.done&&(n=o.return)&&n.call(o)}finally{if(s)throw s.error}}return i}function O(){for(var t=[],e=0;e<arguments.length;e++)t=t.concat(E(arguments[e]));return t}var g=function(){function t(e,n){this.target=n,this.type=e}return t}(),C=function(t){y(e,t);function e(n,o){var r=t.call(this,"error",o)||this;return r.message=n.message,r.error=n,r}return e}(g),T=function(t){y(e,t);function e(n,o,r){n===void 0&&(n=1e3),o===void 0&&(o="");var i=t.call(this,"close",r)||this;return i.wasClean=!0,i.code=n,i.reason=o,i}return e}(g);/*!
 * Reconnecting WebSocket
 * by Pedro Ladaria <pedro.ladaria@gmail.com>
 * https://github.com/pladaria/reconnecting-websocket
 * License MIT
 */var L=function(){if(typeof WebSocket<"u")return WebSocket},S=function(t){return typeof t<"u"&&!!t&&t.CLOSING===2},l={maxReconnectionDelay:1e4,minReconnectionDelay:1e3+Math.random()*4e3,minUptime:5e3,reconnectionDelayGrowFactor:1.3,connectionTimeout:4e3,maxRetries:1/0,maxEnqueuedMessages:1/0,startClosed:!1,debug:!1},x=function(){function t(e,n,o){var r=this;o===void 0&&(o={}),this._listeners={error:[],message:[],open:[],close:[]},this._retryCount=-1,this._shouldReconnect=!0,this._connectLock=!1,this._binaryType="blob",this._closeCalled=!1,this._messageQueue=[],this.onclose=null,this.onerror=null,this.onmessage=null,this.onopen=null,this._handleOpen=function(i){r._debug("open event");var s=r._options.minUptime,c=s===void 0?l.minUptime:s;clearTimeout(r._connectTimeout),r._uptimeTimeout=setTimeout(function(){return r._acceptOpen()},c),r._ws.binaryType=r._binaryType,r._messageQueue.forEach(function(u){return r._ws.send(u)}),r._messageQueue=[],r.onopen&&r.onopen(i),r._listeners.open.forEach(function(u){return r._callEventListener(i,u)})},this._handleMessage=function(i){r._debug("message event"),r.onmessage&&r.onmessage(i),r._listeners.message.forEach(function(s){return r._callEventListener(i,s)})},this._handleError=function(i){r._debug("error event",i.message),r._disconnect(void 0,i.message==="TIMEOUT"?"timeout":void 0),r.onerror&&r.onerror(i),r._debug("exec error listeners"),r._listeners.error.forEach(function(s){return r._callEventListener(i,s)}),r._connect()},this._handleClose=function(i){r._debug("close event"),r._clearTimeouts(),r._shouldReconnect&&r._connect(),r.onclose&&r.onclose(i),r._listeners.close.forEach(function(s){return r._callEventListener(i,s)})},this._url=e,this._protocols=n,this._options=o,this._options.startClosed&&(this._shouldReconnect=!1),this._connect()}return Object.defineProperty(t,"CONNECTING",{get:function(){return 0},enumerable:!0,configurable:!0}),Object.defineProperty(t,"OPEN",{get:function(){return 1},enumerable:!0,configurable:!0}),Object.defineProperty(t,"CLOSING",{get:function(){return 2},enumerable:!0,configurable:!0}),Object.defineProperty(t,"CLOSED",{get:function(){return 3},enumerable:!0,configurable:!0}),Object.defineProperty(t.prototype,"CONNECTING",{get:function(){return t.CONNECTING},enumerable:!0,configurable:!0}),Object.defineProperty(t.prototype,"OPEN",{get:function(){return t.OPEN},enumerable:!0,configurable:!0}),Object.defineProperty(t.prototype,"CLOSING",{get:function(){return t.CLOSING},enumerable:!0,configurable:!0}),Object.defineProperty(t.prototype,"CLOSED",{get:function(){return t.CLOSED},enumerable:!0,configurable:!0}),Object.defineProperty(t.prototype,"binaryType",{get:function(){return this._ws?this._ws.binaryType:this._binaryType},set:function(e){this._binaryType=e,this._ws&&(this._ws.binaryType=e)},enumerable:!0,configurable:!0}),Object.defineProperty(t.prototype,"retryCount",{get:function(){return Math.max(this._retryCount,0)},enumerable:!0,configurable:!0}),Object.defineProperty(t.prototype,"bufferedAmount",{get:function(){var e=this._messageQueue.reduce(function(n,o){return typeof o=="string"?n+=o.length:o instanceof Blob?n+=o.size:n+=o.byteLength,n},0);return e+(this._ws?this._ws.bufferedAmount:0)},enumerable:!0,configurable:!0}),Object.defineProperty(t.prototype,"extensions",{get:function(){return this._ws?this._ws.extensions:""},enumerable:!0,configurable:!0}),Object.defineProperty(t.prototype,"protocol",{get:function(){return this._ws?this._ws.protocol:""},enumerable:!0,configurable:!0}),Object.defineProperty(t.prototype,"readyState",{get:function(){return this._ws?this._ws.readyState:this._options.startClosed?t.CLOSED:t.CONNECTING},enumerable:!0,configurable:!0}),Object.defineProperty(t.prototype,"url",{get:function(){return this._ws?this._ws.url:""},enumerable:!0,configurable:!0}),t.prototype.close=function(e,n){if(e===void 0&&(e=1e3),this._closeCalled=!0,this._shouldReconnect=!1,this._clearTimeouts(),!this._ws){this._debug("close enqueued: no ws instance");return}if(this._ws.readyState===this.CLOSED){this._debug("close: already closed");return}this._ws.close(e,n)},t.prototype.reconnect=function(e,n){this._shouldReconnect=!0,this._closeCalled=!1,this._retryCount=-1,!this._ws||this._ws.readyState===this.CLOSED?this._connect():(this._disconnect(e,n),this._connect())},t.prototype.send=function(e){if(this._ws&&this._ws.readyState===this.OPEN)this._debug("send",e),this._ws.send(e);else{var n=this._options.maxEnqueuedMessages,o=n===void 0?l.maxEnqueuedMessages:n;this._messageQueue.length<o&&(this._debug("enqueue",e),this._messageQueue.push(e))}},t.prototype.addEventListener=function(e,n){this._listeners[e]&&this._listeners[e].push(n)},t.prototype.dispatchEvent=function(e){var n,o,r=this._listeners[e.type];if(r)try{for(var i=w(r),s=i.next();!s.done;s=i.next()){var c=s.value;this._callEventListener(e,c)}}catch(u){n={error:u}}finally{try{s&&!s.done&&(o=i.return)&&o.call(i)}finally{if(n)throw n.error}}return!0},t.prototype.removeEventListener=function(e,n){this._listeners[e]&&(this._listeners[e]=this._listeners[e].filter(function(o){return o!==n}))},t.prototype._debug=function(){for(var e=[],n=0;n<arguments.length;n++)e[n]=arguments[n];this._options.debug&&console.log.apply(console,O(["RWS>"],e))},t.prototype._getNextDelay=function(){var e=this._options,n=e.reconnectionDelayGrowFactor,o=n===void 0?l.reconnectionDelayGrowFactor:n,r=e.minReconnectionDelay,i=r===void 0?l.minReconnectionDelay:r,s=e.maxReconnectionDelay,c=s===void 0?l.maxReconnectionDelay:s,u=0;return this._retryCount>0&&(u=i*Math.pow(o,this._retryCount-1),u>c&&(u=c)),this._debug("next delay",u),u},t.prototype._wait=function(){var e=this;return new Promise(function(n){setTimeout(n,e._getNextDelay())})},t.prototype._getNextUrl=function(e){if(typeof e=="string")return Promise.resolve(e);if(typeof e=="function"){var n=e();if(typeof n=="string")return Promise.resolve(n);if(n.then)return n}throw Error("Invalid URL")},t.prototype._connect=function(){var e=this;if(!(this._connectLock||!this._shouldReconnect)){this._connectLock=!0;var n=this._options,o=n.maxRetries,r=o===void 0?l.maxRetries:o,i=n.connectionTimeout,s=i===void 0?l.connectionTimeout:i,c=n.WebSocket,u=c===void 0?L():c;if(this._retryCount>=r){this._debug("max retries reached",this._retryCount,">=",r);return}if(this._retryCount++,this._debug("connect",this._retryCount),this._removeListeners(),!S(u))throw Error("No valid WebSocket class provided");this._wait().then(function(){return e._getNextUrl(e._url)}).then(function(_){e._closeCalled||(e._debug("connect",{url:_,protocols:e._protocols}),e._ws=e._protocols?new u(_,e._protocols):new u(_),e._ws.binaryType=e._binaryType,e._connectLock=!1,e._addListeners(),e._connectTimeout=setTimeout(function(){return e._handleTimeout()},s))})}},t.prototype._handleTimeout=function(){this._debug("timeout event"),this._handleError(new C(Error("TIMEOUT"),this))},t.prototype._disconnect=function(e,n){if(e===void 0&&(e=1e3),this._clearTimeouts(),!!this._ws){this._removeListeners();try{this._ws.close(e,n),this._handleClose(new T(e,n,this))}catch{}}},t.prototype._acceptOpen=function(){this._debug("accept open"),this._retryCount=0},t.prototype._callEventListener=function(e,n){"handleEvent"in n?n.handleEvent(e):n(e)},t.prototype._removeListeners=function(){this._ws&&(this._debug("removeListeners"),this._ws.removeEventListener("open",this._handleOpen),this._ws.removeEventListener("close",this._handleClose),this._ws.removeEventListener("message",this._handleMessage),this._ws.removeEventListener("error",this._handleError))},t.prototype._addListeners=function(){this._ws&&(this._debug("addListeners"),this._ws.addEventListener("open",this._handleOpen),this._ws.addEventListener("close",this._handleClose),this._ws.addEventListener("message",this._handleMessage),this._ws.addEventListener("error",this._handleError))},t.prototype._clearTimeouts=function(){clearTimeout(this._connectTimeout),clearTimeout(this._uptimeTimeout)},t}();const a=d(null),f=d(new Map),m=d(new Map),h=new Map;h.set("log_change",[]);h.set("status_change",[]);function b(t,e){h.has(t)&&h.get(t).push(e)}function j(t="ws://127.0.0.1:8888/websocket"){return v(()=>{a.value==null&&(a.value=new x(t),a.value.onmessage=N,a.value.onopen=R,a.value.onerror=D,a.value.onclose=P)}),{websocketTarget:a,logMessage:f,statusMessage:m}}const N=t=>{if(t.data=="pong")return;const e=JSON.parse(t.data);e.msgType=="log_change"&&h.get("log_change").map(n=>n(e.data)),e.msgType=="status_change"&&h.get("status_change").map(n=>n(e.data))},R=t=>{var e=setInterval(()=>{a.value.readyState==1?a.value.send("ping"):clearInterval(e)},5e3)},D=t=>{console.log("onError",t)},P=t=>{console.log("onClose",t)};b("log_change",t=>{t.map(e=>{f.value.has(e.id)||f.value.set(e.id,[]),f.value.get(e.id).push(e)})});b("status_change",t=>{m.value.set(t.pid,t.status)});export{b as r,j as u};
