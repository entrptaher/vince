!function(){"use strict";var l=window.location,s=window.document,u=s.currentScript,p=u.getAttribute("data-api")||new URL(u.src).origin+"/api/event";function c(t,e){t&&console.warn("Ignoring Event: "+t),e&&e.callback&&e.callback()}function t(t,e){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(l.hostname)||"file:"===l.protocol)return c("localhost",e);if((window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)&&!window.__plausible)return c(null,e);try{if("true"===window.localStorage.plausible_ignore)return c("localStorage flag",e)}catch(t){}var n=u&&u.getAttribute("data-include"),i=u&&u.getAttribute("data-exclude");if("pageview"===t){n=!n||n.split(",").some(a),i=i&&i.split(",").some(a);if(!n||i)return c("exclusion rule",e)}function a(t){return l.pathname.match(new RegExp("^"+t.trim().replace(/\*\*/g,".*").replace(/([^\.])\*/g,"$1[^\\s/]*")+"/?$"))}var n={},i=(n.n=t,n.u=l.href,n.d=u.getAttribute("data-domain"),n.r=s.referrer||null,e&&e.meta&&(n.m=JSON.stringify(e.meta)),e&&e.props&&(n.p=e.props),e&&e.revenue&&(n.$=e.revenue),u.getAttributeNames().filter(function(t){return"event-"===t.substring(0,6)})),r=n.p||{},o=(i.forEach(function(t){var e=t.replace("event-",""),t=u.getAttribute(t);r[e]=r[e]||t}),n.p=r,new XMLHttpRequest);o.open("POST",p,!0),o.setRequestHeader("Content-Type","text/plain"),o.send(JSON.stringify(n)),o.onreadystatechange=function(){4===o.readyState&&e&&e.callback&&e.callback({status:o.status})}}var e=window.plausible&&window.plausible.q||[];window.plausible=t;for(var n,i=0;i<e.length;i++)t.apply(this,e[i]);function a(){n!==l.pathname&&(n=l.pathname,t("pageview"))}var r,o=window.history;o.pushState&&(r=o.pushState,o.pushState=function(){r.apply(this,arguments),a()},window.addEventListener("popstate",a)),"prerender"===s.visibilityState?s.addEventListener("visibilitychange",function(){n||"visible"!==s.visibilityState||a()}):a()}();