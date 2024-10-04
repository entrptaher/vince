!function(){"use strict";var p=window.location,l=window.document,o=l.getElementById("plausible"),s=o.getAttribute("data-api")||(g=(g=o).src.split("/"),u=g[0],g=g[2],u+"//"+g+"/api/event");function c(t,e){t&&console.warn("Ignoring Event: "+t),e&&e.callback&&e.callback()}function t(t,e){try{if("true"===window.localStorage.plausible_ignore)return c("localStorage flag",e)}catch(t){}var a=o&&o.getAttribute("data-include"),i=o&&o.getAttribute("data-exclude");if("pageview"===t){a=!a||a.split(",").some(n),i=i&&i.split(",").some(n);if(!a||i)return c("exclusion rule",e)}function n(t){return p.pathname.match(new RegExp("^"+t.trim().replace(/\*\*/g,".*").replace(/([^\.])\*/g,"$1[^\\s/]*")+"/?$"))}var a={},r=(a.n=t,a.u=p.href,a.d=o.getAttribute("data-domain"),a.r=l.referrer||null,e&&e.meta&&(a.m=JSON.stringify(e.meta)),e&&e.props&&(a.p=e.props),new XMLHttpRequest);r.open("POST",s,!0),r.setRequestHeader("Content-Type","text/plain"),r.send(JSON.stringify(a)),r.onreadystatechange=function(){4===r.readyState&&e&&e.callback&&e.callback({status:r.status})}}var e=window.plausible&&window.plausible.q||[];window.plausible=t;for(var a,i=0;i<e.length;i++)t.apply(this,e[i]);function n(){a!==p.pathname&&(a=p.pathname,t("pageview"))}var r,u=window.history;u.pushState&&(r=u.pushState,u.pushState=function(){r.apply(this,arguments),n()},window.addEventListener("popstate",n)),"prerender"===l.visibilityState?l.addEventListener("visibilitychange",function(){a||"visible"!==l.visibilityState||n()}):n();var d=1;function f(t){var e,a,i,n,r,p,l;function o(){n||(n=!0,window.location=i.href)}"auxclick"===t.type&&t.button!==d||(e=function(t){for(;t&&(void 0===t.tagName||!(e=t)||!e.tagName||"a"!==e.tagName.toLowerCase()||!t.href);)t=t.parentNode;var e;return t}(t.target),a=e&&e.href&&e.href.split("?")[0],(p=a)&&(l=p.split(".").pop(),w.some(function(t){return t===l}))&&(n=!(p={name:"File Download",props:{url:a}}),!function(t,e){if(!t.defaultPrevented)return e=!e.target||e.target.match(/^_(self|parent|top)$/i),t=!(t.ctrlKey||t.metaKey||t.shiftKey)&&"click"===t.type,e&&t}(a=t,i=e)?(r={props:p.props},plausible(p.name,r)):(r={props:p.props,callback:o},plausible(p.name,r),setTimeout(o,5e3),a.preventDefault())))}l.addEventListener("click",f),l.addEventListener("auxclick",f);var g=["pdf","xlsx","docx","txt","rtf","csv","exe","key","pps","ppt","pptx","7z","pkg","rar","gz","zip","avi","mov","mp4","mpeg","wmv","midi","mp3","wav","wma","dmg"],m=o.getAttribute("file-types"),v=o.getAttribute("add-file-types"),w=m&&m.split(",")||v&&v.split(",").concat(g)||g}();