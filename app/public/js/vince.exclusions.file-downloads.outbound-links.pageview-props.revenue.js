!function(){"use strict";var p=window.location,l=window.document,s=l.currentScript,u=s.getAttribute("data-api")||new URL(s.src).origin+"/api/event";function c(e,t){e&&console.warn("Ignoring Event: "+e),t&&t.callback&&t.callback()}function e(e,t){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(p.hostname)||"file:"===p.protocol)return c("localhost",t);if((window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)&&!window.__plausible)return c(null,t);try{if("true"===window.localStorage.plausible_ignore)return c("localStorage flag",t)}catch(e){}var n=s&&s.getAttribute("data-include"),a=s&&s.getAttribute("data-exclude");if("pageview"===e){n=!n||n.split(",").some(i),a=a&&a.split(",").some(i);if(!n||a)return c("exclusion rule",t)}function i(e){return p.pathname.match(new RegExp("^"+e.trim().replace(/\*\*/g,".*").replace(/([^\.])\*/g,"$1[^\\s/]*")+"/?$"))}var n={},a=(n.n=e,n.u=p.href,n.d=s.getAttribute("data-domain"),n.r=l.referrer||null,t&&t.meta&&(n.m=JSON.stringify(t.meta)),t&&t.props&&(n.p=t.props),t&&t.revenue&&(n.$=t.revenue),s.getAttributeNames().filter(function(e){return"event-"===e.substring(0,6)})),r=n.p||{},o=(a.forEach(function(e){var t=e.replace("event-",""),e=s.getAttribute(e);r[t]=r[t]||e}),n.p=r,new XMLHttpRequest);o.open("POST",u,!0),o.setRequestHeader("Content-Type","text/plain"),o.send(JSON.stringify(n)),o.onreadystatechange=function(){4===o.readyState&&t&&t.callback&&t.callback({status:o.status})}}var t=window.plausible&&window.plausible.q||[];window.plausible=e;for(var n,a=0;a<t.length;a++)e.apply(this,t[a]);function i(){n!==p.pathname&&(n=p.pathname,e("pageview"))}var r,o=window.history;o.pushState&&(r=o.pushState,o.pushState=function(){r.apply(this,arguments),i()},window.addEventListener("popstate",i)),"prerender"===l.visibilityState?l.addEventListener("visibilitychange",function(){n||"visible"!==l.visibilityState||i()}):i();var d=1;function f(e){var t,n,a,i;if("auxclick"!==e.type||e.button===d)return t=function(e){for(;e&&(void 0===e.tagName||!(t=e)||!t.tagName||"a"!==t.tagName.toLowerCase()||!e.href);)e=e.parentNode;var t;return e}(e.target),n=t&&t.href&&t.href.split("?")[0],(a=t)&&a.href&&a.host&&a.host!==p.host?v(e,t,{name:"Outbound Link: Click",props:{url:t.href}}):(a=n)&&(i=a.split(".").pop(),m.some(function(e){return e===i}))?v(e,t,{name:"File Download",props:{url:n}}):void 0}function v(e,t,n){var a,i=!1;function r(){i||(i=!0,window.location=t.href)}!function(e,t){if(!e.defaultPrevented)return t=!t.target||t.target.match(/^_(self|parent|top)$/i),e=!(e.ctrlKey||e.metaKey||e.shiftKey)&&"click"===e.type,t&&e}(e,t)?((a={props:n.props}).revenue=n.revenue,plausible(n.name,a)):((a={props:n.props,callback:r}).revenue=n.revenue,plausible(n.name,a),setTimeout(r,5e3),e.preventDefault())}l.addEventListener("click",f),l.addEventListener("auxclick",f);var o=["pdf","xlsx","docx","txt","rtf","csv","exe","key","pps","ppt","pptx","7z","pkg","rar","gz","zip","avi","mov","mp4","mpeg","wmv","midi","mp3","wav","wma","dmg"],w=s.getAttribute("file-types"),g=s.getAttribute("add-file-types"),m=w&&w.split(",")||g&&g.split(",").concat(o)||o}();