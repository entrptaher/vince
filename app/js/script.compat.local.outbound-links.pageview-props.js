!function(){"use strict";var t,l=window.location,o=window.document,s=o.getElementById("plausible"),p=s.getAttribute("data-api")||(t=(t=s).src.split("/"),c=t[0],t=t[2],c+"//"+t+"/api/event");function e(t,e){try{if("true"===window.localStorage.plausible_ignore)return n=e,(a="localStorage flag")&&console.warn("Ignoring Event: "+a),void(n&&n.callback&&n.callback())}catch(t){}var a={},n=(a.n=t,a.u=l.href,a.d=s.getAttribute("data-domain"),a.r=o.referrer||null,e&&e.meta&&(a.m=JSON.stringify(e.meta)),e&&e.props&&(a.p=e.props),s.getAttributeNames().filter(function(t){return"event-"===t.substring(0,6)})),i=a.p||{},r=(n.forEach(function(t){var e=t.replace("event-",""),t=s.getAttribute(t);i[e]=i[e]||t}),a.p=i,new XMLHttpRequest);r.open("POST",p,!0),r.setRequestHeader("Content-Type","text/plain"),r.send(JSON.stringify(a)),r.onreadystatechange=function(){4===r.readyState&&e&&e.callback&&e.callback({status:r.status})}}var a=window.plausible&&window.plausible.q||[];window.plausible=e;for(var n,i=0;i<a.length;i++)e.apply(this,a[i]);function r(){n!==l.pathname&&(n=l.pathname,e("pageview"))}var u,c=window.history;c.pushState&&(u=c.pushState,c.pushState=function(){u.apply(this,arguments),r()},window.addEventListener("popstate",r)),"prerender"===o.visibilityState?o.addEventListener("visibilitychange",function(){n||"visible"!==o.visibilityState||r()}):r();var d=1;function f(t){var e,a,n,i,r;function o(){n||(n=!0,window.location=a.href)}"auxclick"===t.type&&t.button!==d||((e=function(t){for(;t&&(void 0===t.tagName||!(e=t)||!e.tagName||"a"!==e.tagName.toLowerCase()||!t.href);)t=t.parentNode;var e;return t}(t.target))&&e.href&&e.href.split("?")[0],(r=e)&&r.href&&r.host&&r.host!==l.host&&(r=t,t={name:"Outbound Link: Click",props:{url:(a=e).href}},n=!1,!function(t,e){if(!t.defaultPrevented)return e=!e.target||e.target.match(/^_(self|parent|top)$/i),t=!(t.ctrlKey||t.metaKey||t.shiftKey)&&"click"===t.type,e&&t}(r,a)?(i={props:t.props},plausible(t.name,i)):(i={props:t.props,callback:o},plausible(t.name,i),setTimeout(o,5e3),r.preventDefault())))}o.addEventListener("click",f),o.addEventListener("auxclick",f)}();