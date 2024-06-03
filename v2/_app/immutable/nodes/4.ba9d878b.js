import{r as G,v as J,q as K,w as E}from"../chunks/pages.db7e4f8c.js";import{e as Q}from"../chunks/index.0bc363c4.js";import{s as O,g as U,c as D,h as V,i as W,d as k,k as X,l as R,m as C,T as Y}from"../chunks/scheduler.559b666f.js";import{S as N,i as H,d as g,v as T,j as v,k as h,e as p,a as _,m as $,n as B,b as f,h as Z,t as m,c as x,l as w,s as ee}from"../chunks/index.59aa5467.js";import{A as y,C as q,P as M}from"../chunks/PreviousNextPage.84f993bd.js";import{A as z}from"../chunks/ArgsList.75e4c884.js";const te=G(),ne=({params:n,url:e})=>{const i=J(n.directive),r=K(e.pathname);if(!i||!r)throw Q(404,`Directive ${n.directive} not found.`);return{directive:i,page:r}},fe=Object.freeze(Object.defineProperty({__proto__:null,load:ne,prerender:te},Symbol.toStringTag,{value:"Module"})),ie="src/lib/components/directive/DirectiveDetails.svelte";function I(n){let e=n[0].name+"",i;const r={c:function(){i=R(e)},l:function(o){i=C(o,e)},m:function(o,u){_(o,i,u)},p:function(o,u){u&1&&e!==(e=o[0].name+"")&&ee(i,e)},d:function(o){o&&p(i)}};return g("SvelteRegisterBlock",{block:r,id:I.name,type:"slot",source:"(12:2) <AnchorHeader id={'title'} depth={1}>",ctx:n}),r}function P(n){let e,i,r,c;e=new y({props:{id:"arguments",depth:2,$$slots:{default:[L]},$$scope:{ctx:n}},$$inline:!0}),r=new z({props:{data:n[1]},$$inline:!0});const o={c:function(){v(e.$$.fragment),i=D(),v(r.$$.fragment)},l:function(t){h(e.$$.fragment,t),i=k(t),h(r.$$.fragment,t)},m:function(t,s){$(e,t,s),_(t,i,s),$(r,t,s),c=!0},p:function(t,s){const a={};s&4&&(a.$$scope={dirty:s,ctx:t}),e.$set(a);const l={};s&2&&(l.data=t[1]),r.$set(l)},i:function(t){c||(f(e.$$.fragment,t),f(r.$$.fragment,t),c=!0)},o:function(t){m(e.$$.fragment,t),m(r.$$.fragment,t),c=!1},d:function(t){t&&p(i),w(e,t),w(r,t)}};return g("SvelteRegisterBlock",{block:o,id:P.name,type:"if",source:"(18:2) {#if allowedArgs.length > 0}",ctx:n}),o}function L(n){let e;const i={c:function(){e=R("Arguments")},l:function(c){e=C(c,"Arguments")},m:function(c,o){_(c,e,o)},d:function(c){c&&p(e)}};return g("SvelteRegisterBlock",{block:i,id:L.name,type:"slot",source:"(19:4) <AnchorHeader id={'arguments'} depth={2}>",ctx:n}),i}function A(n){let e,i,r,c,o,u;i=new y({props:{id:"title",depth:1,$$slots:{default:[I]},$$scope:{ctx:n}},$$inline:!0}),c=new q({props:{source:n[0].description},$$inline:!0});let t=n[1].length>0&&P(n);const s={c:function(){e=U("section"),v(i.$$.fragment),r=D(),v(c.$$.fragment),o=D(),t&&t.c(),this.h()},l:function(l){e=V(l,"SECTION",{});var d=W(e);h(i.$$.fragment,d),r=k(d),h(c.$$.fragment,d),o=k(d),t&&t.l(d),d.forEach(p),this.h()},h:function(){X(e,ie,13,0,372)},m:function(l,d){_(l,e,d),$(i,e,null),B(e,r),$(c,e,null),B(e,o),t&&t.m(e,null),u=!0},p:function(l,[d]){const b={};d&5&&(b.$$scope={dirty:d,ctx:l}),i.$set(b);const S={};d&1&&(S.source=l[0].description),c.$set(S),l[1].length>0?t?(t.p(l,d),d&2&&f(t,1)):(t=P(l),t.c(),f(t,1),t.m(e,null)):t&&(Z(),m(t,1,1,()=>{t=null}),x())},i:function(l){u||(f(i.$$.fragment,l),f(c.$$.fragment,l),f(t),u=!0)},o:function(l){m(i.$$.fragment,l),m(c.$$.fragment,l),m(t),u=!1},d:function(l){l&&p(e),w(i),w(c),t&&t.d()}};return g("SvelteRegisterBlock",{block:s,id:A.name,type:"component",source:"",ctx:n}),s}function oe(n,e,i){let{$$slots:r={},$$scope:c}=e;T("DirectiveDetails",r,[]);let{directive:o}=e,u;n.$$.on_mount.push(function(){o===void 0&&!("directive"in e||n.$$.bound[n.$$.props.directive])&&console.warn("<DirectiveDetails> was created without expected prop 'directive'")});const t=["directive"];return Object.keys(e).forEach(s=>{!~t.indexOf(s)&&s.slice(0,2)!=="$$"&&s!=="slot"&&console.warn(`<DirectiveDetails> was created with unknown prop '${s}'`)}),n.$$set=s=>{"directive"in s&&i(0,o=s.directive)},n.$capture_state=()=>({getAllowedArgumentsByDirective:E,AnchorHeader:y,CarbonMarkdown:q,ArgsList:z,directive:o,allowedArgs:u}),n.$inject_state=s=>{"directive"in s&&i(0,o=s.directive),"allowedArgs"in s&&i(1,u=s.allowedArgs)},e&&"$$inject"in e&&n.$inject_state(e.$$inject),n.$$.update=()=>{n.$$.dirty&1&&i(1,u=E(o))},[o,u]}class F extends N{constructor(e){super(e),H(this,e,oe,A,O,{directive:0}),g("SvelteRegisterComponent",{component:this,tagName:"DirectiveDetails",options:e,id:A.name})}get directive(){throw new Error("<DirectiveDetails>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'")}set directive(e){throw new Error("<DirectiveDetails>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'")}}function j(n){let e,i,r,c,o,u;document.title=e="Directive - "+n[0].directive.name,r=new F({props:{directive:n[0].directive},$$inline:!0}),o=new M({props:{page:n[0].page},$$inline:!0});const t={c:function(){i=D(),v(r.$$.fragment),c=D(),v(o.$$.fragment)},l:function(a){Y("svelte-fycjkl",document.head).forEach(p),i=k(a),h(r.$$.fragment,a),c=k(a),h(o.$$.fragment,a)},m:function(a,l){_(a,i,l),$(r,a,l),_(a,c,l),$(o,a,l),u=!0},p:function(a,[l]){(!u||l&1)&&e!==(e="Directive - "+a[0].directive.name)&&(document.title=e);const d={};l&1&&(d.directive=a[0].directive),r.$set(d);const b={};l&1&&(b.page=a[0].page),o.$set(b)},i:function(a){u||(f(r.$$.fragment,a),f(o.$$.fragment,a),u=!0)},o:function(a){m(r.$$.fragment,a),m(o.$$.fragment,a),u=!1},d:function(a){a&&(p(i),p(c)),w(r,a),w(o,a)}};return g("SvelteRegisterBlock",{block:t,id:j.name,type:"component",source:"",ctx:n}),t}function re(n,e,i){let{$$slots:r={},$$scope:c}=e;T("Page",r,[]);let{data:o}=e;n.$$.on_mount.push(function(){o===void 0&&!("data"in e||n.$$.bound[n.$$.props.data])&&console.warn("<Page> was created without expected prop 'data'")});const u=["data"];return Object.keys(e).forEach(t=>{!~u.indexOf(t)&&t.slice(0,2)!=="$$"&&t!=="slot"&&console.warn(`<Page> was created with unknown prop '${t}'`)}),n.$$set=t=>{"data"in t&&i(0,o=t.data)},n.$capture_state=()=>({PreviousNextPage:M,DirectiveDetails:F,data:o}),n.$inject_state=t=>{"data"in t&&i(0,o=t.data)},e&&"$$inject"in e&&n.$inject_state(e.$$inject),[o]}class me extends N{constructor(e){super(e),H(this,e,re,j,O,{data:0}),g("SvelteRegisterComponent",{component:this,tagName:"Page",options:e,id:j.name})}get data(){throw new Error("<Page>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'")}set data(e){throw new Error("<Page>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'")}}export{me as component,fe as universal};
