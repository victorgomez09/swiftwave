import{z as q,A as O,q as B}from"../chunks/pages.c01f7e9b.js";import{e as N}from"../chunks/index.0bc363c4.js";import{s as R,c as g,T,d as _}from"../chunks/scheduler.559b666f.js";import{S as k,i as x,d as j,v as z,j as h,e as l,k as v,a as w,m as y,b as $,t as P,l as b}from"../chunks/index.59aa5467.js";import{F as Q,Q as S}from"../chunks/FieldDetails.09821b36.js";import{P as E}from"../chunks/PreviousNextPage.da4340d1.js";const C=q(),D=({params:n,url:t})=>{const i=O(n.query),o=B(t.pathname);if(!i||!o)throw N(404,`Query ${n.query} not found.`);return{field:i,page:o}},I=Object.freeze(Object.defineProperty({__proto__:null,load:D,prerender:C},Symbol.toStringTag,{value:"Module"}));function f(n){let t,i,o,d,a,c;document.title=t="Query - "+n[0].field.name,o=new Q({props:{field:n[0].field,type:S.QUERY},$$inline:!0}),a=new E({props:{page:n[0].page},$$inline:!0});const s={c:function(){i=g(),h(o.$$.fragment),d=g(),h(a.$$.fragment)},l:function(e){T("svelte-19yffit",document.head).forEach(l),i=_(e),v(o.$$.fragment,e),d=_(e),v(a.$$.fragment,e)},m:function(e,r){w(e,i,r),y(o,e,r),w(e,d,r),y(a,e,r),c=!0},p:function(e,[r]){(!c||r&1)&&t!==(t="Query - "+e[0].field.name)&&(document.title=t);const m={};r&1&&(m.field=e[0].field),o.$set(m);const p={};r&1&&(p.page=e[0].page),a.$set(p)},i:function(e){c||($(o.$$.fragment,e),$(a.$$.fragment,e),c=!0)},o:function(e){P(o.$$.fragment,e),P(a.$$.fragment,e),c=!1},d:function(e){e&&(l(i),l(d)),b(o,e),b(a,e)}};return j("SvelteRegisterBlock",{block:s,id:f.name,type:"component",source:"",ctx:n}),s}function F(n,t,i){let{$$slots:o={},$$scope:d}=t;z("Page",o,[]);let{data:a}=t;n.$$.on_mount.push(function(){a===void 0&&!("data"in t||n.$$.bound[n.$$.props.data])&&console.warn("<Page> was created without expected prop 'data'")});const c=["data"];return Object.keys(t).forEach(s=>{!~c.indexOf(s)&&s.slice(0,2)!=="$$"&&s!=="slot"&&console.warn(`<Page> was created with unknown prop '${s}'`)}),n.$$set=s=>{"data"in s&&i(0,a=s.data)},n.$capture_state=()=>({QueryType:S,FieldDetails:Q,PreviousNextPage:E,data:a}),n.$inject_state=s=>{"data"in s&&i(0,a=s.data)},t&&"$$inject"in t&&n.$inject_state(t.$$inject),[a]}class J extends k{constructor(t){super(t),x(this,t,F,f,R,{data:0}),j("SvelteRegisterComponent",{component:this,tagName:"Page",options:t,id:f.name})}get data(){throw new Error("<Page>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'")}set data(t){throw new Error("<Page>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'")}}export{J as component,I as universal};
