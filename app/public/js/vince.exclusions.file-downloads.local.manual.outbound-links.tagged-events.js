!function(){"use strict";var o=window.location,p=window.document,l=p.currentScript,u=l.getAttribute("data-api")||new URL(l.src).origin+"/api/event";function s(e,t){e&&console.warn("Ignoring Event: "+e),t&&t.callback&&t.callback()}function e(e,t){try{if("true"===window.localStorage.plausible_ignore)return s("localStorage flag",t)}catch(e){}var a=l&&l.getAttribute("data-include"),r=l&&l.getAttribute("data-exclude");if("pageview"===e){a=!a||a.split(",").some(n),r=r&&r.split(",").some(n);if(!a||r)return s("exclusion rule",t)}function n(e){return o.pathname.match(new RegExp("^"+e.trim().replace(/\*\*/g,".*").replace(/([^\.])\*/g,"$1[^\\s/]*")+"/?$"))}var a={},i=(a.n=e,a.u=t&&t.u?t.u:o.href,a.d=l.getAttribute("data-domain"),a.r=p.referrer||null,t&&t.meta&&(a.m=JSON.stringify(t.meta)),t&&t.props&&(a.p=t.props),new XMLHttpRequest);i.open("POST",u,!0),i.setRequestHeader("Content-Type","text/plain"),i.send(JSON.stringify(a)),i.onreadystatechange=function(){4===i.readyState&&t&&t.callback&&t.callback({status:i.status})}}var t=window.plausible&&window.plausible.q||[];window.plausible=e;for(var a=0;a<t.length;a++)e.apply(this,t[a]);function c(e){return e&&e.tagName&&"a"===e.tagName.toLowerCase()}var f=1;function r(e){if("auxclick"!==e.type||e.button===f){var t,a,r=function(e){for(;e&&(void 0===e.tagName||!c(e)||!e.href);)e=e.parentNode;return e}(e.target),n=r&&r.href&&r.href.split("?")[0];if(!function e(t,a){if(!t||b<a)return!1;if(h(t))return!0;return e(t.parentNode,a+1)}(r,0))return(t=r)&&t.href&&t.host&&t.host!==o.host?d(e,r,{name:"Outbound Link: Click",props:{url:r.href}}):(t=n)&&(a=t.split(".").pop(),g.some(function(e){return e===a}))?d(e,r,{name:"File Download",props:{url:n}}):void 0}}function d(e,t,a){var r,n=!1;function i(){n||(n=!0,window.location=t.href)}!function(e,t){if(!e.defaultPrevented)return t=!t.target||t.target.match(/^_(self|parent|top)$/i),e=!(e.ctrlKey||e.metaKey||e.shiftKey)&&"click"===e.type,t&&e}(e,t)?(r={props:a.props},plausible(a.name,r)):(r={props:a.props,callback:i},plausible(a.name,r),setTimeout(i,5e3),e.preventDefault())}p.addEventListener("click",r),p.addEventListener("auxclick",r);var n=["pdf","xlsx","docx","txt","rtf","csv","exe","key","pps","ppt","pptx","7z","pkg","rar","gz","zip","avi","mov","mp4","mpeg","wmv","midi","mp3","wav","wma","dmg"],i=l.getAttribute("file-types"),m=l.getAttribute("add-file-types"),g=i&&i.split(",")||m&&m.split(",").concat(n)||n;function v(e){var e=h(e)?e:e&&e.parentNode,t={name:null,props:{}},a=e&&e.classList;if(a)for(var r=0;r<a.length;r++){var n,i=a.item(r).match(/plausible-event-(.+)(=|--)(.+)/);i&&(n=i[1],i=i[3].replace(/\+/g," "),"name"==n.toLowerCase()?t.name=i:t.props[n]=i)}return t}var b=3;function w(e){if("auxclick"!==e.type||e.button===f){for(var t,a,r,n,i=e.target,o=0;o<=b&&i;o++){if((r=i)&&r.tagName&&"form"===r.tagName.toLowerCase())return;c(i)&&(t=i),h(i)&&(a=i),i=i.parentNode}a&&(n=v(a),t?(n.props.url=t.href,d(e,t,n)):((e={}).props=n.props,plausible(n.name,e)))}}function h(e){var t=e&&e.classList;if(t)for(var a=0;a<t.length;a++)if(t.item(a).match(/plausible-event-name(=|--)(.+)/))return!0;return!1}p.addEventListener("submit",function(e){var t,a=e.target,r=v(a);function n(){t||(t=!0,a.submit())}r.name&&(e.preventDefault(),t=!1,setTimeout(n,5e3),e={props:r.props,callback:n},plausible(r.name,e))}),p.addEventListener("click",w),p.addEventListener("auxclick",w)}();