(this["webpackJsonpchakra-ui-typescript"]=this["webpackJsonpchakra-ui-typescript"]||[]).push([[0],{66:function(e,t,n){e.exports=n(67)},67:function(e,t,n){"use strict";n.r(t);var a=n(21),l=n(63),c=n(0),r=n.n(c),o=n(15),i=n.n(o),s=n(4),u=n(64),m=n(41),E=n(40),p=n(42),d=n(12),f=n(107),b=n(104),h=n(65),L=n(47),O=n(105),A=n(106),v=n(31),_=n(43),j=n(14),C=n(48),B=n(72),g=B.BACKEND_URL,N=B.FEED_LABEL,T=B.APP_TITLE,y=B.COMMENT_LABEL,S=B.COMMENT_BUTTON_LABEL,w=B.COMMENT_UNAVAILABLE_LABEL;B.FEED_UNAVAILABLE_LABEL;function D(e){var t=e.title,n=e.desc,o=e.id,i=Object(l.a)(e,["title","desc","id"]),d=Object(c.useState)(null),f=Object(a.a)(d,2),b=(f[0],f[1]),L=Object(c.useState)(null),O=Object(a.a)(L,2),A=O[0],_=O[1];function j(){fetch(g+"/comments?event_id="+o,{method:"GET"}).then((function(e){return e.json()})).then((function(e){_(e)}),(function(e){b(e)}))}return Object(c.useEffect)((function(){return j()}),[]),r.a.createElement(s.a,Object.assign({p:5,shadow:"md",borderWidth:"1px"},i),r.a.createElement(u.a,{fontSize:"xl"},t),r.a.createElement(m.a,{mt:4,mb:4},n),r.a.createElement(v.a,{allowMultiple:!0},r.a.createElement(v.d,null,r.a.createElement(v.b,{_expanded:{bg:"#1A365D",color:"white"}},r.a.createElement(s.a,{flex:"1",textAlign:"left"},y," (",null===A?0:A.length,")"),r.a.createElement(v.c,null)),r.a.createElement(v.e,null,r.a.createElement(h.a,{spacing:3,align:"center"},null===A?r.a.createElement(m.a,null,w):A.map((function(e){return r.a.createElement(s.a,{p:2,w:"100%",borderWidth:"1px"},r.a.createElement(m.a,null,e.content))}))),r.a.createElement("form",{style:{width:"100%"},onSubmit:function(e){e.preventDefault();var t=new FormData(e.target);fetch(g+"/comments",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({commenter:"supratik_das",content:t.get("content"),event_id:o})}).then((function(e){return e.json()})).then((function(e){"success"==e.status&&j()}))}},r.a.createElement(C.a,{mt:5},r.a.createElement(E.a,{type:"text",id:"content",name:"content","aria-describedby":"text-helper-text"}),r.a.createElement(p.a,{variantColor:"teal",variant:"solid",type:"submit",mt:2},S)))))))}function M(e){var t=Object(b.a)(),n=t.isOpen,a=t.onOpen,l=t.onClose,c=Object(A.a)(),o=r.a.useRef();return r.a.createElement(r.a.Fragment,null,r.a.createElement(p.a,{ref:o,variantColor:"teal",onClick:a},"Create an ",N),r.a.createElement(_.a,{isOpen:n,placement:"top",onClose:l,finalFocusRef:o},r.a.createElement(_.d,null),r.a.createElement(_.c,null,r.a.createElement("form",{style:{width:"100%"},onSubmit:function(t){t.preventDefault();var n=new FormData(t.target);fetch(g+"/events",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({name:n.get("name"),description:n.get("description")})}).then((function(e){return e.json()})).then((function(t){"success"===t.status?(e.callBack(!0),c({title:"Event created.",description:"",status:"success",duration:5e3,isClosable:!0})):c({title:"Event creation failed",description:"",status:"error",duration:5e3,isClosable:!0})}))}},r.a.createElement(_.b,null),r.a.createElement(j.f,null,"Create an ",N),r.a.createElement(j.b,null,r.a.createElement(E.a,{name:"name",placeholder:"Event Name..."}),r.a.createElement(O.a,{name:"description",mt:5,placeholder:"Event Description..."})),r.a.createElement(j.e,null,r.a.createElement(p.a,{variant:"outline",mr:3,onClick:l},"Cancel"),r.a.createElement(p.a,{type:"submit",color:"blue"},"Save"))))))}function x(){var e=Object(c.useState)(null),t=Object(a.a)(e,2),n=(t[0],t[1]),l=Object(c.useState)(null),o=Object(a.a)(l,2),i=o[0],E=o[1],p=Object(c.useState)(!1),b=Object(a.a)(p,2),O=(b[0],b[1]);return Object(c.useEffect)((function(){return fetch(g+"/events",{method:"GET"}).then((function(e){return e.json()})).then((function(e){E(e)}),(function(e){n(e)}))}),[]),r.a.createElement(d.a,null,r.a.createElement(f.a,null),r.a.createElement(s.a,{bg:"#1A365D",w:"100%",p:4,color:"white"},r.a.createElement(L.a,{align:"center",justify:"space-between"},r.a.createElement(u.a,{as:"h2",size:"xl"},T),r.a.createElement(M,{callBack:O}))),r.a.createElement(h.a,{pl:20,pr:20,pt:10,spacing:8},null===i?r.a.createElement(m.a,null,"Loading"):i.map((function(e){return r.a.createElement(D,{title:e.name,desc:e.description,id:e.id})}))))}var k=document.getElementById("root");i.a.render(r.a.createElement(x,null),k)},72:function(e){e.exports=JSON.parse('{"BACKEND_URL":"http://localhost:8080/api/v1","FEED_LABEL":"Event","APP_TITLE":"Your online events catalogue","COMMENT_LABEL":"Comments","COMMENT_BUTTON_LABEL":"Comment","COMMENT_UNAVAILABLE_LABEL":"No comments on this event yet.","FEED_UNAVAILABLE_LABEL":"No events are available yet."}')}},[[66,1,2]]]);
//# sourceMappingURL=main.700216f9.chunk.js.map