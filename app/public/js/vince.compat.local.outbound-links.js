!function(){"use strict";var t,l=window.location,r=window.document,o=r.getElementById("plausible"),s=o.getAttribute("data-api")||(t=(t=o).src.split("/"),u=t[0],t=t[2],u+"//"+t+"/api/event");function e(t,e){try{if("true"===window.localStorage.plausible_ignore)return a=e,(n="localStorage flag")&&console.warn("Ignoring Event: "+n),void(a&&a.callback&&a.callback())}catch(t){}var a,n={},i=(n.n=t,n.u=l.href,n.d=o.getAttribute("data-domain"),n.r=r.referrer||null,e&&e.meta&&(n.m=JSON.stringify(e.meta)),e&&e.props&&(n.p=e.props),new XMLHttpRequest);i.open("POST",s,!0),i.setRequestHeader("Content-Type","text/plain"),i.send(JSON.stringify(n)),i.onreadystatechange=function(){4===i.readyState&&e&&e.callback&&e.callback({status:i.status})}}var a=window.plausible&&window.plausible.q||[];window.plausible=e;for(var n,i=0;i<a.length;i++)e.apply(this,a[i]);function p(){n!==l.pathname&&(n=l.pathname,e("pageview"))}var c,u=window.history;u.pushState&&(c=u.pushState,u.pushState=function(){c.apply(this,arguments),p()},window.addEventListener("popstate",p)),"prerender"===r.visibilityState?r.addEventListener("visibilitychange",function(){n||"visible"!==r.visibilityState||p()}):p();var d=1;function f(t){var e,a,n,i,r;function o(){n||(n=!0,window.location=a.href)}"auxclick"===t.type&&t.button!==d||((e=function(t){for(;t&&(void 0===t.tagName||!(e=t)||!e.tagName||"a"!==e.tagName.toLowerCase()||!t.href);)t=t.parentNode;var e;return t}(t.target))&&e.href&&e.href.split("?")[0],(r=e)&&r.href&&r.host&&r.host!==l.host&&(r=t,t={name:"Outbound Link: Click",props:{url:(a=e).href}},n=!1,!function(t,e){if(!t.defaultPrevented)return e=!e.target||e.target.match(/^_(self|parent|top)$/i),t=!(t.ctrlKey||t.metaKey||t.shiftKey)&&"click"===t.type,e&&t}(r,a)?(i={props:t.props},plausible(t.name,i)):(i={props:t.props,callback:o},plausible(t.name,i),setTimeout(o,5e3),r.preventDefault())))}r.addEventListener("click",f),r.addEventListener("auxclick",f)}();