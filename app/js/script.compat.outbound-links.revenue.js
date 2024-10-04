!function(){"use strict";var e,l=window.location,i=window.document,r=i.getElementById("plausible"),o=r.getAttribute("data-api")||(e=(e=r).src.split("/"),d=e[0],e=e[2],d+"//"+e+"/api/event");function s(e,t){e&&console.warn("Ignoring Event: "+e),t&&t.callback&&t.callback()}function t(e,t){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(l.hostname)||"file:"===l.protocol)return s("localhost",t);if((window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)&&!window.__plausible)return s(null,t);try{if("true"===window.localStorage.plausible_ignore)return s("localStorage flag",t)}catch(e){}var a={},n=(a.n=e,a.u=l.href,a.d=r.getAttribute("data-domain"),a.r=i.referrer||null,t&&t.meta&&(a.m=JSON.stringify(t.meta)),t&&t.props&&(a.p=t.props),t&&t.revenue&&(a.$=t.revenue),new XMLHttpRequest);n.open("POST",o,!0),n.setRequestHeader("Content-Type","text/plain"),n.send(JSON.stringify(a)),n.onreadystatechange=function(){4===n.readyState&&t&&t.callback&&t.callback({status:n.status})}}var a=window.plausible&&window.plausible.q||[];window.plausible=t;for(var n,u=0;u<a.length;u++)t.apply(this,a[u]);function p(){n!==l.pathname&&(n=l.pathname,t("pageview"))}var c,d=window.history;d.pushState&&(c=d.pushState,d.pushState=function(){c.apply(this,arguments),p()},window.addEventListener("popstate",p)),"prerender"===i.visibilityState?i.addEventListener("visibilitychange",function(){n||"visible"!==i.visibilityState||p()}):p();var f=1;function w(e){var t,a,n,i,r;function o(){n||(n=!0,window.location=a.href)}"auxclick"===e.type&&e.button!==f||((t=function(e){for(;e&&(void 0===e.tagName||!(t=e)||!t.tagName||"a"!==t.tagName.toLowerCase()||!e.href);)e=e.parentNode;var t;return e}(e.target))&&t.href&&t.href.split("?")[0],(r=t)&&r.href&&r.host&&r.host!==l.host&&(r=e,e={name:"Outbound Link: Click",props:{url:(a=t).href}},n=!1,!function(e,t){if(!e.defaultPrevented)return t=!t.target||t.target.match(/^_(self|parent|top)$/i),e=!(e.ctrlKey||e.metaKey||e.shiftKey)&&"click"===e.type,t&&e}(r,a)?((i={props:e.props}).revenue=e.revenue,plausible(e.name,i)):((i={props:e.props,callback:o}).revenue=e.revenue,plausible(e.name,i),setTimeout(o,5e3),r.preventDefault())))}i.addEventListener("click",w),i.addEventListener("auxclick",w)}();