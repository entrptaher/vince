!function(){"use strict";var o=window.location,l=window.document,s=l.currentScript,c=s.getAttribute("data-api")||new URL(s.src).origin+"/api/event";function t(t,e){try{if("true"===window.localStorage.plausible_ignore)return i=e,(n="localStorage flag")&&console.warn("Ignoring Event: "+n),void(i&&i.callback&&i.callback())}catch(t){}var n={},i=(n.n=t,n.u=o.href,n.d=s.getAttribute("data-domain"),n.r=l.referrer||null,e&&e.meta&&(n.m=JSON.stringify(e.meta)),e&&e.props&&(n.p=e.props),s.getAttributeNames().filter(function(t){return"event-"===t.substring(0,6)})),a=n.p||{},r=(i.forEach(function(t){var e=t.replace("event-",""),t=s.getAttribute(t);a[e]=a[e]||t}),n.p=a,n.h=1,new XMLHttpRequest);r.open("POST",c,!0),r.setRequestHeader("Content-Type","text/plain"),r.send(JSON.stringify(n)),r.onreadystatechange=function(){4===r.readyState&&e&&e.callback&&e.callback({status:r.status})}}var e=window.plausible&&window.plausible.q||[];window.plausible=t;for(var n,i=0;i<e.length;i++)t.apply(this,e[i]);function a(){n=o.pathname,t("pageview")}window.addEventListener("hashchange",a),"prerender"===l.visibilityState?l.addEventListener("visibilitychange",function(){n||"visible"!==l.visibilityState||a()}):a()}();