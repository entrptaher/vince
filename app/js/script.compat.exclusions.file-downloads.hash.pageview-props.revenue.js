!function(){"use strict";var o=window.location,p=window.document,s=p.getElementById("plausible"),u=s.getAttribute("data-api")||(f=(f=s).src.split("/"),l=f[0],f=f[2],l+"//"+f+"/api/event");function c(e,t){e&&console.warn("Ignoring Event: "+e),t&&t.callback&&t.callback()}function e(e,t){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(o.hostname)||"file:"===o.protocol)return c("localhost",t);if((window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)&&!window.__plausible)return c(null,t);try{if("true"===window.localStorage.plausible_ignore)return c("localStorage flag",t)}catch(e){}var a=s&&s.getAttribute("data-include"),n=s&&s.getAttribute("data-exclude");if("pageview"===e){a=!a||a.split(",").some(i),n=n&&n.split(",").some(i);if(!a||n)return c("exclusion rule",t)}function i(e){var t=o.pathname;return(t+=o.hash).match(new RegExp("^"+e.trim().replace(/\*\*/g,".*").replace(/([^\.])\*/g,"$1[^\\s/]*")+"/?$"))}var a={},n=(a.n=e,a.u=o.href,a.d=s.getAttribute("data-domain"),a.r=p.referrer||null,t&&t.meta&&(a.m=JSON.stringify(t.meta)),t&&t.props&&(a.p=t.props),t&&t.revenue&&(a.$=t.revenue),s.getAttributeNames().filter(function(e){return"event-"===e.substring(0,6)})),r=a.p||{},l=(n.forEach(function(e){var t=e.replace("event-",""),e=s.getAttribute(e);r[t]=r[t]||e}),a.p=r,a.h=1,new XMLHttpRequest);l.open("POST",u,!0),l.setRequestHeader("Content-Type","text/plain"),l.send(JSON.stringify(a)),l.onreadystatechange=function(){4===l.readyState&&t&&t.callback&&t.callback({status:l.status})}}var t=window.plausible&&window.plausible.q||[];window.plausible=e;for(var a,n=0;n<t.length;n++)e.apply(this,t[n]);function i(){a=o.pathname,e("pageview")}window.addEventListener("hashchange",i),"prerender"===p.visibilityState?p.addEventListener("visibilitychange",function(){a||"visible"!==p.visibilityState||i()}):i();var d=1;function r(e){var t,a,n,i,r,l,o;function p(){i||(i=!0,window.location=n.href)}"auxclick"===e.type&&e.button!==d||(t=function(e){for(;e&&(void 0===e.tagName||!(t=e)||!t.tagName||"a"!==t.tagName.toLowerCase()||!e.href);)e=e.parentNode;var t;return e}(e.target),a=t&&t.href&&t.href.split("?")[0],(l=a)&&(o=l.split(".").pop(),g.some(function(e){return e===o}))&&(i=!(l={name:"File Download",props:{url:a}}),!function(e,t){if(!e.defaultPrevented)return t=!t.target||t.target.match(/^_(self|parent|top)$/i),e=!(e.ctrlKey||e.metaKey||e.shiftKey)&&"click"===e.type,t&&e}(a=e,n=t)?((r={props:l.props}).revenue=l.revenue,plausible(l.name,r)):((r={props:l.props,callback:p}).revenue=l.revenue,plausible(l.name,r),setTimeout(p,5e3),a.preventDefault())))}p.addEventListener("click",r),p.addEventListener("auxclick",r);var l=["pdf","xlsx","docx","txt","rtf","csv","exe","key","pps","ppt","pptx","7z","pkg","rar","gz","zip","avi","mov","mp4","mpeg","wmv","midi","mp3","wav","wma","dmg"],f=s.getAttribute("file-types"),v=s.getAttribute("add-file-types"),g=f&&f.split(",")||v&&v.split(",").concat(l)||l}();