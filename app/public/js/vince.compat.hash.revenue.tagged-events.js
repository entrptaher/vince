!function(){"use strict";var e,t,r=window.location,i=window.document,o=i.getElementById("plausible"),l=o.getAttribute("data-api")||(e=(e=o).src.split("/"),t=e[0],e=e[2],t+"//"+e+"/api/event");function u(e,t){e&&console.warn("Ignoring Event: "+e),t&&t.callback&&t.callback()}function n(e,t){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(r.hostname)||"file:"===r.protocol)return u("localhost",t);if((window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)&&!window.__plausible)return u(null,t);try{if("true"===window.localStorage.plausible_ignore)return u("localStorage flag",t)}catch(e){}var n={},a=(n.n=e,n.u=r.href,n.d=o.getAttribute("data-domain"),n.r=i.referrer||null,t&&t.meta&&(n.m=JSON.stringify(t.meta)),t&&t.props&&(n.p=t.props),t&&t.revenue&&(n.$=t.revenue),n.h=1,new XMLHttpRequest);a.open("POST",l,!0),a.setRequestHeader("Content-Type","text/plain"),a.send(JSON.stringify(n)),a.onreadystatechange=function(){4===a.readyState&&t&&t.callback&&t.callback({status:a.status})}}var a=window.plausible&&window.plausible.q||[];window.plausible=n;for(var s,c=0;c<a.length;c++)n.apply(this,a[c]);function p(){s=r.pathname,n("pageview")}function f(e){return e&&e.tagName&&"a"===e.tagName.toLowerCase()}window.addEventListener("hashchange",p),"prerender"===i.visibilityState?i.addEventListener("visibilitychange",function(){s||"visible"!==i.visibilityState||p()}):p();var v=1;function d(e){"auxclick"===e.type&&e.button!==v||((e=function(e){for(;e&&(void 0===e.tagName||!f(e)||!e.href);)e=e.parentNode;return e}(e.target))&&e.href&&e.href.split("?")[0],function e(t,n){if(!t||g<n)return!1;if(h(t))return!0;return e(t.parentNode,n+1)}(e,0))}function m(e,t,n){var a,r=!1;function i(){r||(r=!0,window.location=t.href)}!function(e,t){if(!e.defaultPrevented)return t=!t.target||t.target.match(/^_(self|parent|top)$/i),e=!(e.ctrlKey||e.metaKey||e.shiftKey)&&"click"===e.type,t&&e}(e,t)?((a={props:n.props}).revenue=n.revenue,plausible(n.name,a)):((a={props:n.props,callback:i}).revenue=n.revenue,plausible(n.name,a),setTimeout(i,5e3),e.preventDefault())}function w(e){var e=h(e)?e:e&&e.parentNode,t={name:null,props:{},revenue:{}},n=e&&e.classList;if(n)for(var a=0;a<n.length;a++){var r,i,o=n.item(a),l=o.match(/plausible-event-(.+)(=|--)(.+)/),l=(l&&(r=l[1],i=l[3].replace(/\+/g," "),"name"==r.toLowerCase()?t.name=i:t.props[r]=i),o.match(/plausible-revenue-(.+)(=|--)(.+)/));l&&(r=l[1],i=l[3],t.revenue[r]=i)}return t}i.addEventListener("click",d),i.addEventListener("auxclick",d);var g=3;function b(e){if("auxclick"!==e.type||e.button===v){for(var t,n,a,r,i=e.target,o=0;o<=g&&i;o++){if((a=i)&&a.tagName&&"form"===a.tagName.toLowerCase())return;f(i)&&(t=i),h(i)&&(n=i),i=i.parentNode}n&&(r=w(n),t?(r.props.url=t.href,m(e,t,r)):((e={}).props=r.props,e.revenue=r.revenue,plausible(r.name,e)))}}function h(e){var t=e&&e.classList;if(t)for(var n=0;n<t.length;n++)if(t.item(n).match(/plausible-event-name(=|--)(.+)/))return!0;return!1}i.addEventListener("submit",function(e){var t,n=e.target,a=w(n);function r(){t||(t=!0,n.submit())}a.name&&(e.preventDefault(),t=!1,setTimeout(r,5e3),(e={props:a.props,callback:r}).revenue=a.revenue,plausible(a.name,e))}),i.addEventListener("click",b),i.addEventListener("auxclick",b)}();