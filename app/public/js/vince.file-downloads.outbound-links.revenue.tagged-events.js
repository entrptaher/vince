!function(){"use strict";var i=window.location,a=window.document,o=a.currentScript,p=o.getAttribute("data-api")||new URL(o.src).origin+"/api/event";function u(e,t){e&&console.warn("Ignoring Event: "+e),t&&t.callback&&t.callback()}function e(e,t){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(i.hostname)||"file:"===i.protocol)return u("localhost",t);if((window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)&&!window.__plausible)return u(null,t);try{if("true"===window.localStorage.plausible_ignore)return u("localStorage flag",t)}catch(e){}var n={},r=(n.n=e,n.u=i.href,n.d=o.getAttribute("data-domain"),n.r=a.referrer||null,t&&t.meta&&(n.m=JSON.stringify(t.meta)),t&&t.props&&(n.p=t.props),t&&t.revenue&&(n.$=t.revenue),new XMLHttpRequest);r.open("POST",p,!0),r.setRequestHeader("Content-Type","text/plain"),r.send(JSON.stringify(n)),r.onreadystatechange=function(){4===r.readyState&&t&&t.callback&&t.callback({status:r.status})}}var t=window.plausible&&window.plausible.q||[];window.plausible=e;for(var n,r=0;r<t.length;r++)e.apply(this,t[r]);function l(){n!==i.pathname&&(n=i.pathname,e("pageview"))}var s,c=window.history;function f(e){return e&&e.tagName&&"a"===e.tagName.toLowerCase()}c.pushState&&(s=c.pushState,c.pushState=function(){s.apply(this,arguments),l()},window.addEventListener("popstate",l)),"prerender"===a.visibilityState?a.addEventListener("visibilitychange",function(){n||"visible"!==a.visibilityState||l()}):l();var v=1;function d(e){if("auxclick"!==e.type||e.button===v){var t,n,r=function(e){for(;e&&(void 0===e.tagName||!f(e)||!e.href);)e=e.parentNode;return e}(e.target),a=r&&r.href&&r.href.split("?")[0];if(!function e(t,n){if(!t||y<n)return!1;if(L(t))return!0;return e(t.parentNode,n+1)}(r,0))return(t=r)&&t.href&&t.host&&t.host!==i.host?m(e,r,{name:"Outbound Link: Click",props:{url:r.href}}):(t=a)&&(n=t.split(".").pop(),h.some(function(e){return e===n}))?m(e,r,{name:"File Download",props:{url:a}}):void 0}}function m(e,t,n){var r,a=!1;function i(){a||(a=!0,window.location=t.href)}!function(e,t){if(!e.defaultPrevented)return t=!t.target||t.target.match(/^_(self|parent|top)$/i),e=!(e.ctrlKey||e.metaKey||e.shiftKey)&&"click"===e.type,t&&e}(e,t)?((r={props:n.props}).revenue=n.revenue,plausible(n.name,r)):((r={props:n.props,callback:i}).revenue=n.revenue,plausible(n.name,r),setTimeout(i,5e3),e.preventDefault())}a.addEventListener("click",d),a.addEventListener("auxclick",d);var c=["pdf","xlsx","docx","txt","rtf","csv","exe","key","pps","ppt","pptx","7z","pkg","rar","gz","zip","avi","mov","mp4","mpeg","wmv","midi","mp3","wav","wma","dmg"],w=o.getAttribute("file-types"),g=o.getAttribute("add-file-types"),h=w&&w.split(",")||g&&g.split(",").concat(c)||c;function b(e){var e=L(e)?e:e&&e.parentNode,t={name:null,props:{},revenue:{}},n=e&&e.classList;if(n)for(var r=0;r<n.length;r++){var a,i,o=n.item(r),p=o.match(/plausible-event-(.+)(=|--)(.+)/),p=(p&&(a=p[1],i=p[3].replace(/\+/g," "),"name"==a.toLowerCase()?t.name=i:t.props[a]=i),o.match(/plausible-revenue-(.+)(=|--)(.+)/));p&&(a=p[1],i=p[3],t.revenue[a]=i)}return t}var y=3;function k(e){if("auxclick"!==e.type||e.button===v){for(var t,n,r,a,i=e.target,o=0;o<=y&&i;o++){if((r=i)&&r.tagName&&"form"===r.tagName.toLowerCase())return;f(i)&&(t=i),L(i)&&(n=i),i=i.parentNode}n&&(a=b(n),t?(a.props.url=t.href,m(e,t,a)):((e={}).props=a.props,e.revenue=a.revenue,plausible(a.name,e)))}}function L(e){var t=e&&e.classList;if(t)for(var n=0;n<t.length;n++)if(t.item(n).match(/plausible-event-name(=|--)(.+)/))return!0;return!1}a.addEventListener("submit",function(e){var t,n=e.target,r=b(n);function a(){t||(t=!0,n.submit())}r.name&&(e.preventDefault(),t=!1,setTimeout(a,5e3),(e={props:r.props,callback:a}).revenue=r.revenue,plausible(r.name,e))}),a.addEventListener("click",k),a.addEventListener("auxclick",k)}();