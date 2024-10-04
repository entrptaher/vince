!function(){"use strict";var o=window.location,l=window.document,p=l.currentScript,s=p.getAttribute("data-api")||new URL(p.src).origin+"/api/event";function u(e,t){e&&console.warn("Ignoring Event: "+e),t&&t.callback&&t.callback()}function e(e,t){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(o.hostname)||"file:"===o.protocol)return u("localhost",t);if((window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)&&!window.__plausible)return u(null,t);try{if("true"===window.localStorage.plausible_ignore)return u("localStorage flag",t)}catch(e){}var a=p&&p.getAttribute("data-include"),n=p&&p.getAttribute("data-exclude");if("pageview"===e){a=!a||a.split(",").some(i),n=n&&n.split(",").some(i);if(!a||n)return u("exclusion rule",t)}function i(e){var t=o.pathname;return(t+=o.hash).match(new RegExp("^"+e.trim().replace(/\*\*/g,".*").replace(/([^\.])\*/g,"$1[^\\s/]*")+"/?$"))}var a={},r=(a.n=e,a.u=o.href,a.d=p.getAttribute("data-domain"),a.r=l.referrer||null,t&&t.meta&&(a.m=JSON.stringify(t.meta)),t&&t.props&&(a.p=t.props),t&&t.revenue&&(a.$=t.revenue),a.h=1,new XMLHttpRequest);r.open("POST",s,!0),r.setRequestHeader("Content-Type","text/plain"),r.send(JSON.stringify(a)),r.onreadystatechange=function(){4===r.readyState&&t&&t.callback&&t.callback({status:r.status})}}var t=window.plausible&&window.plausible.q||[];window.plausible=e;for(var a,n=0;n<t.length;n++)e.apply(this,t[n]);function i(){a=o.pathname,e("pageview")}window.addEventListener("hashchange",i),"prerender"===l.visibilityState?l.addEventListener("visibilitychange",function(){a||"visible"!==l.visibilityState||i()}):i();var r=1;function c(e){var t,a,n,i;if("auxclick"!==e.type||e.button===r)return t=function(e){for(;e&&(void 0===e.tagName||!(t=e)||!t.tagName||"a"!==t.tagName.toLowerCase()||!e.href);)e=e.parentNode;var t;return e}(e.target),a=t&&t.href&&t.href.split("?")[0],(n=t)&&n.href&&n.host&&n.host!==o.host?d(e,t,{name:"Outbound Link: Click",props:{url:t.href}}):(n=a)&&(i=n.split(".").pop(),g.some(function(e){return e===i}))?d(e,t,{name:"File Download",props:{url:a}}):void 0}function d(e,t,a){var n,i=!1;function r(){i||(i=!0,window.location=t.href)}!function(e,t){if(!e.defaultPrevented)return t=!t.target||t.target.match(/^_(self|parent|top)$/i),e=!(e.ctrlKey||e.metaKey||e.shiftKey)&&"click"===e.type,t&&e}(e,t)?((n={props:a.props}).revenue=a.revenue,plausible(a.name,n)):((n={props:a.props,callback:r}).revenue=a.revenue,plausible(a.name,n),setTimeout(r,5e3),e.preventDefault())}l.addEventListener("click",c),l.addEventListener("auxclick",c);var f=["pdf","xlsx","docx","txt","rtf","csv","exe","key","pps","ppt","pptx","7z","pkg","rar","gz","zip","avi","mov","mp4","mpeg","wmv","midi","mp3","wav","wma","dmg"],v=p.getAttribute("file-types"),w=p.getAttribute("add-file-types"),g=v&&v.split(",")||w&&w.split(",").concat(f)||f}();