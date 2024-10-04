!function(){"use strict";var o=window.location,s=window.document,u=s.currentScript,l=u.getAttribute("data-api")||new URL(u.src).origin+"/api/event";function c(e,t){e&&console.warn("Ignoring Event: "+e),t&&t.callback&&t.callback()}function e(e,t){try{if("true"===window.localStorage.plausible_ignore)return c("localStorage flag",t)}catch(e){}var a=u&&u.getAttribute("data-include"),n=u&&u.getAttribute("data-exclude");if("pageview"===e){a=!a||a.split(",").some(r),n=n&&n.split(",").some(r);if(!a||n)return c("exclusion rule",t)}function r(e){return o.pathname.match(new RegExp("^"+e.trim().replace(/\*\*/g,".*").replace(/([^\.])\*/g,"$1[^\\s/]*")+"/?$"))}var a={},n=(a.n=e,a.u=o.href,a.d=u.getAttribute("data-domain"),a.r=s.referrer||null,t&&t.meta&&(a.m=JSON.stringify(t.meta)),t&&t.props&&(a.p=t.props),u.getAttributeNames().filter(function(e){return"event-"===e.substring(0,6)})),i=a.p||{},p=(n.forEach(function(e){var t=e.replace("event-",""),e=u.getAttribute(e);i[t]=i[t]||e}),a.p=i,new XMLHttpRequest);p.open("POST",l,!0),p.setRequestHeader("Content-Type","text/plain"),p.send(JSON.stringify(a)),p.onreadystatechange=function(){4===p.readyState&&t&&t.callback&&t.callback({status:p.status})}}var t=window.plausible&&window.plausible.q||[];window.plausible=e;for(var a,n=0;n<t.length;n++)e.apply(this,t[n]);function r(){a!==o.pathname&&(a=o.pathname,e("pageview"))}var i,p=window.history;function f(e){return e&&e.tagName&&"a"===e.tagName.toLowerCase()}p.pushState&&(i=p.pushState,p.pushState=function(){i.apply(this,arguments),r()},window.addEventListener("popstate",r)),"prerender"===s.visibilityState?s.addEventListener("visibilitychange",function(){a||"visible"!==s.visibilityState||r()}):r();var d=1;function m(e){var t,a,n,r;if("auxclick"!==e.type||e.button===d)return t=function(e){for(;e&&(void 0===e.tagName||!f(e)||!e.href);)e=e.parentNode;return e}(e.target),a=t&&t.href&&t.href.split("?")[0],!function e(t,a){if(!t||y<a)return!1;if(x(t))return!0;return e(t.parentNode,a+1)}(t,0)&&(n=a)&&(r=n.split(".").pop(),w.some(function(e){return e===r}))?v(e,t,{name:"File Download",props:{url:a}}):void 0}function v(e,t,a){var n,r=!1;function i(){r||(r=!0,window.location=t.href)}!function(e,t){if(!e.defaultPrevented)return t=!t.target||t.target.match(/^_(self|parent|top)$/i),e=!(e.ctrlKey||e.metaKey||e.shiftKey)&&"click"===e.type,t&&e}(e,t)?(n={props:a.props},plausible(a.name,n)):(n={props:a.props,callback:i},plausible(a.name,n),setTimeout(i,5e3),e.preventDefault())}s.addEventListener("click",m),s.addEventListener("auxclick",m);var p=["pdf","xlsx","docx","txt","rtf","csv","exe","key","pps","ppt","pptx","7z","pkg","rar","gz","zip","avi","mov","mp4","mpeg","wmv","midi","mp3","wav","wma","dmg"],g=u.getAttribute("file-types"),b=u.getAttribute("add-file-types"),w=g&&g.split(",")||b&&b.split(",").concat(p)||p;function h(e){var e=x(e)?e:e&&e.parentNode,t={name:null,props:{}},a=e&&e.classList;if(a)for(var n=0;n<a.length;n++){var r,i=a.item(n).match(/plausible-event-(.+)(=|--)(.+)/);i&&(r=i[1],i=i[3].replace(/\+/g," "),"name"==r.toLowerCase()?t.name=i:t.props[r]=i)}return t}var y=3;function k(e){if("auxclick"!==e.type||e.button===d){for(var t,a,n,r,i=e.target,p=0;p<=y&&i;p++){if((n=i)&&n.tagName&&"form"===n.tagName.toLowerCase())return;f(i)&&(t=i),x(i)&&(a=i),i=i.parentNode}a&&(r=h(a),t?(r.props.url=t.href,v(e,t,r)):((e={}).props=r.props,plausible(r.name,e)))}}function x(e){var t=e&&e.classList;if(t)for(var a=0;a<t.length;a++)if(t.item(a).match(/plausible-event-name(=|--)(.+)/))return!0;return!1}s.addEventListener("submit",function(e){var t,a=e.target,n=h(a);function r(){t||(t=!0,a.submit())}n.name&&(e.preventDefault(),t=!1,setTimeout(r,5e3),e={props:n.props,callback:r},plausible(n.name,e))}),s.addEventListener("click",k),s.addEventListener("auxclick",k)}();