!function(){"use strict";var e,t,o=window.location,l=window.document,s=l.getElementById("plausible"),c=s.getAttribute("data-api")||(e=(e=s).src.split("/"),t=e[0],e=e[2],t+"//"+e+"/api/event");function p(e,t){e&&console.warn("Ignoring Event: "+e),t&&t.callback&&t.callback()}function n(e,t){try{if("true"===window.localStorage.plausible_ignore)return p("localStorage flag",t)}catch(e){}var n=s&&s.getAttribute("data-include"),r=s&&s.getAttribute("data-exclude");if("pageview"===e){n=!n||n.split(",").some(a),r=r&&r.split(",").some(a);if(!n||r)return p("exclusion rule",t)}function a(e){var t=o.pathname;return(t+=o.hash).match(new RegExp("^"+e.trim().replace(/\*\*/g,".*").replace(/([^\.])\*/g,"$1[^\\s/]*")+"/?$"))}var n={},r=(n.n=e,n.u=o.href,n.d=s.getAttribute("data-domain"),n.r=l.referrer||null,t&&t.meta&&(n.m=JSON.stringify(t.meta)),t&&t.props&&(n.p=t.props),t&&t.revenue&&(n.$=t.revenue),s.getAttributeNames().filter(function(e){return"event-"===e.substring(0,6)})),i=n.p||{},u=(r.forEach(function(e){var t=e.replace("event-",""),e=s.getAttribute(e);i[t]=i[t]||e}),n.p=i,n.h=1,new XMLHttpRequest);u.open("POST",c,!0),u.setRequestHeader("Content-Type","text/plain"),u.send(JSON.stringify(n)),u.onreadystatechange=function(){4===u.readyState&&t&&t.callback&&t.callback({status:u.status})}}var r=window.plausible&&window.plausible.q||[];window.plausible=n;for(var a,i=0;i<r.length;i++)n.apply(this,r[i]);function u(){a=o.pathname,n("pageview")}function f(e){return e&&e.tagName&&"a"===e.tagName.toLowerCase()}window.addEventListener("hashchange",u),"prerender"===l.visibilityState?l.addEventListener("visibilitychange",function(){a||"visible"!==l.visibilityState||u()}):u();var v=1;function d(e){var t,n;if("auxclick"!==e.type||e.button===v)return(t=function(e){for(;e&&(void 0===e.tagName||!f(e)||!e.href);)e=e.parentNode;return e}(e.target))&&t.href&&t.href.split("?")[0],!function e(t,n){if(!t||b<n)return!1;if(w(t))return!0;return e(t.parentNode,n+1)}(t,0)&&(n=t)&&n.href&&n.host&&n.host!==o.host?m(e,t,{name:"Outbound Link: Click",props:{url:t.href}}):void 0}function m(e,t,n){var r,a=!1;function i(){a||(a=!0,window.location=t.href)}!function(e,t){if(!e.defaultPrevented)return t=!t.target||t.target.match(/^_(self|parent|top)$/i),e=!(e.ctrlKey||e.metaKey||e.shiftKey)&&"click"===e.type,t&&e}(e,t)?((r={props:n.props}).revenue=n.revenue,plausible(n.name,r)):((r={props:n.props,callback:i}).revenue=n.revenue,plausible(n.name,r),setTimeout(i,5e3),e.preventDefault())}function g(e){var e=w(e)?e:e&&e.parentNode,t={name:null,props:{},revenue:{}},n=e&&e.classList;if(n)for(var r=0;r<n.length;r++){var a,i,u=n.item(r),o=u.match(/plausible-event-(.+)(=|--)(.+)/),o=(o&&(a=o[1],i=o[3].replace(/\+/g," "),"name"==a.toLowerCase()?t.name=i:t.props[a]=i),u.match(/plausible-revenue-(.+)(=|--)(.+)/));o&&(a=o[1],i=o[3],t.revenue[a]=i)}return t}l.addEventListener("click",d),l.addEventListener("auxclick",d);var b=3;function h(e){if("auxclick"!==e.type||e.button===v){for(var t,n,r,a,i=e.target,u=0;u<=b&&i;u++){if((r=i)&&r.tagName&&"form"===r.tagName.toLowerCase())return;f(i)&&(t=i),w(i)&&(n=i),i=i.parentNode}n&&(a=g(n),t?(a.props.url=t.href,m(e,t,a)):((e={}).props=a.props,e.revenue=a.revenue,plausible(a.name,e)))}}function w(e){var t=e&&e.classList;if(t)for(var n=0;n<t.length;n++)if(t.item(n).match(/plausible-event-name(=|--)(.+)/))return!0;return!1}l.addEventListener("submit",function(e){var t,n=e.target,r=g(n);function a(){t||(t=!0,n.submit())}r.name&&(e.preventDefault(),t=!1,setTimeout(a,5e3),(e={props:r.props,callback:a}).revenue=r.revenue,plausible(r.name,e))}),l.addEventListener("click",h),l.addEventListener("auxclick",h)}();