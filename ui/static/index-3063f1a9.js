import{v as Ue,al as Ae,am as Fe,an as W,ao as re,ap as ge,t as q,x as ie,aq as ue,Z as J,r as $,w as I,H as de,ar as le,C as Me,c as P,d as L,aa as ye,G as T,ab as je,Y as he,A as ce,as as Le,at as Ie,z as G,o as Ce,ae as qe,f as t,a4 as Y,au as Ke,av as Ge,aw as Ye,ag as H,a0 as He,$ as oe,ax as We,ay as Ve,az as me,ad as Je,aA as Qe,aB as Xe,aC as pe,V as Ze,aD as ea,S as K,a8 as aa,T as Ee,ak as ta,aE as la,s as Q,_ as oa,J as na,K as Te,L as ve,M as U,g as i,j as A,N as fe,O as be,h as sa,i as ra,P as ne,l as ke,m as we,n as Ne,E as $e,Q as ia,I as ua,aF as da,aG as ca,R as ma,e as se}from"./index-45490d46.js";import{E as pa}from"./el-pagination-e9fd92a9.js";import{a as xe,E as Se}from"./el-table-column-63ea15fa.js";import{c as F,E as va,a as fa}from"./el-select-4076887c.js";import"./el-checkbox-43c79793.js";import{E as ba,a as _a}from"./el-col-439bec59.js";import{u as ga}from"./page-56987b87.js";import{E as ya,a as ha}from"./el-radio-59976266.js";import{Y as Ca}from"./index-d025be05.js";import{v as Va}from"./directive-5b5f6efd.js";import{E as Ea}from"./index-a0985bb2.js";import"./index-cf7261e3.js";const Ta=(l,d,h)=>Ae(l.subTree).filter(a=>{var p;return Fe(a)&&((p=a.type)==null?void 0:p.name)===d&&!!a.component}).map(a=>a.component.uid).map(a=>h[a]).filter(a=>!!a),ka=(l,d)=>{const h={},v=Ue([]);return{children:v,addChild:p=>{h[p.uid]=p,v.value=Ta(l,d,h)},removeChild:p=>{delete h[p],v.value=v.value.filter(b=>b.uid!==p)}}},X=Symbol("tabsRootContextKey"),wa=W({tabs:{type:re(Array),default:()=>ge([])}}),Pe="ElTabBar",Na=q({name:Pe}),$a=q({...Na,props:wa,setup(l,{expose:d}){const h=l,v=ce(),r=ie(X);r||ue(Pe,"<el-tabs><el-tab-bar /></el-tabs>");const a=J("tabs"),p=$(),b=$(),f=()=>{let u=0,E=0;const V=["top","bottom"].includes(r.props.tabPosition)?"width":"height",C=V==="width"?"x":"y",s=C==="x"?"left":"top";return h.tabs.every(k=>{var x,B;const N=(B=(x=v.parent)==null?void 0:x.refs)==null?void 0:B[`tab-${k.uid}`];if(!N)return!1;if(!k.active)return!0;u=N[`offset${F(s)}`],E=N[`client${F(V)}`];const R=window.getComputedStyle(N);return V==="width"&&(E-=Number.parseFloat(R.paddingLeft)+Number.parseFloat(R.paddingRight),u+=Number.parseFloat(R.paddingLeft)),!1}),{[V]:`${E}px`,transform:`translate${F(C)}(${u}px)`}},w=()=>b.value=f(),_=[],o=()=>{var u;_.forEach(V=>V.stop()),_.length=0;const E=(u=v.parent)==null?void 0:u.refs;if(E){for(const V in E)if(V.startsWith("tab-")){const C=E[V];C&&_.push(le(C,w))}}};I(()=>h.tabs,async()=>{await de(),w(),o()},{immediate:!0});const m=le(p,()=>w());return Me(()=>{_.forEach(u=>u.stop()),_.length=0,m.stop()}),d({ref:p,update:w}),(u,E)=>(P(),L("div",{ref_key:"barRef",ref:p,class:ye([T(a).e("active-bar"),T(a).is(T(r).props.tabPosition)]),style:je(b.value)},null,6))}});var xa=he($a,[["__file","tab-bar.vue"]]);const Sa=W({panes:{type:re(Array),default:()=>ge([])},currentName:{type:[String,Number],default:""},editable:Boolean,type:{type:String,values:["card","border-card",""],default:""},stretch:Boolean}),Pa={tabClick:(l,d,h)=>h instanceof Event,tabRemove:(l,d)=>d instanceof Event},_e="ElTabNav",za=q({name:_e,props:Sa,emits:Pa,setup(l,{expose:d,emit:h}){const v=ie(X);v||ue(_e,"<el-tabs><tab-nav /></el-tabs>");const r=J("tabs"),a=Le(),p=Ie(),b=$(),f=$(),w=$(),_=$(),o=$(!1),m=$(0),u=$(!1),E=$(!0),V=G(()=>["top","bottom"].includes(v.props.tabPosition)?"width":"height"),C=G(()=>({transform:`translate${V.value==="width"?"X":"Y"}(-${m.value}px)`})),s=()=>{if(!b.value)return;const c=b.value[`offset${F(V.value)}`],e=m.value;if(!e)return;const n=e>c?e-c:0;m.value=n},k=()=>{if(!b.value||!f.value)return;const c=f.value[`offset${F(V.value)}`],e=b.value[`offset${F(V.value)}`],n=m.value;if(c-n<=e)return;const g=c-n>e*2?n+e:c-e;m.value=g},x=async()=>{const c=f.value;if(!o.value||!w.value||!b.value||!c)return;await de();const e=w.value.querySelector(".is-active");if(!e)return;const n=b.value,g=["top","bottom"].includes(v.props.tabPosition),z=e.getBoundingClientRect(),y=n.getBoundingClientRect(),O=g?c.offsetWidth-y.width:c.offsetHeight-y.height,D=m.value;let S=D;g?(z.left<y.left&&(S=D-(y.left-z.left)),z.right>y.right&&(S=D+z.right-y.right)):(z.top<y.top&&(S=D-(y.top-z.top)),z.bottom>y.bottom&&(S=D+(z.bottom-y.bottom))),S=Math.max(S,0),m.value=Math.min(S,O)},B=()=>{var c;if(!f.value||!b.value)return;l.stretch&&((c=_.value)==null||c.update());const e=f.value[`offset${F(V.value)}`],n=b.value[`offset${F(V.value)}`],g=m.value;n<e?(o.value=o.value||{},o.value.prev=g,o.value.next=g+n<e,e-g<n&&(m.value=e-n)):(o.value=!1,g>0&&(m.value=0))},N=c=>{const e=c.code,{up:n,down:g,left:z,right:y}=H;if(![n,g,z,y].includes(e))return;const O=Array.from(c.currentTarget.querySelectorAll("[role=tab]:not(.is-disabled)")),D=O.indexOf(c.target);let S;e===z||e===n?D===0?S=O.length-1:S=D-1:D<O.length-1?S=D+1:S=0,O[S].focus({preventScroll:!0}),O[S].click(),R()},R=()=>{E.value&&(u.value=!0)},M=()=>u.value=!1;return I(a,c=>{c==="hidden"?E.value=!1:c==="visible"&&setTimeout(()=>E.value=!0,50)}),I(p,c=>{c?setTimeout(()=>E.value=!0,50):E.value=!1}),le(w,B),Ce(()=>setTimeout(()=>x(),0)),qe(()=>B()),d({scrollToActiveTab:x,removeFocus:M}),()=>{const c=o.value?[t("span",{class:[r.e("nav-prev"),r.is("disabled",!o.value.prev)],onClick:s},[t(Y,null,{default:()=>[t(Ke,null,null)]})]),t("span",{class:[r.e("nav-next"),r.is("disabled",!o.value.next)],onClick:k},[t(Y,null,{default:()=>[t(Ge,null,null)]})])]:null,e=l.panes.map((n,g)=>{var z,y,O,D;const S=n.uid,Z=n.props.disabled,ee=(y=(z=n.props.name)!=null?z:n.index)!=null?y:`${g}`,ae=!Z&&(n.isClosable||l.editable);n.index=`${g}`;const Re=ae?t(Y,{class:"is-icon-close",onClick:j=>h("tabRemove",n,j)},{default:()=>[t(Ye,null,null)]}):null,De=((D=(O=n.slots).label)==null?void 0:D.call(O))||n.props.label,Oe=!Z&&n.active?0:-1;return t("div",{ref:`tab-${S}`,class:[r.e("item"),r.is(v.props.tabPosition),r.is("active",n.active),r.is("disabled",Z),r.is("closable",ae),r.is("focus",u.value)],id:`tab-${ee}`,key:`tab-${S}`,"aria-controls":`pane-${ee}`,role:"tab","aria-selected":n.active,tabindex:Oe,onFocus:()=>R(),onBlur:()=>M(),onClick:j=>{M(),h("tabClick",n,ee,j)},onKeydown:j=>{ae&&(j.code===H.delete||j.code===H.backspace)&&h("tabRemove",n,j)}},[De,Re])});return t("div",{ref:w,class:[r.e("nav-wrap"),r.is("scrollable",!!o.value),r.is(v.props.tabPosition)]},[c,t("div",{class:r.e("nav-scroll"),ref:b},[t("div",{class:[r.e("nav"),r.is(v.props.tabPosition),r.is("stretch",l.stretch&&["top","bottom"].includes(v.props.tabPosition))],ref:f,style:C.value,role:"tablist",onKeydown:N},[l.type?null:t(xa,{ref:_,tabs:[...l.panes]},null),e])])])}}}),Ba=W({type:{type:String,values:["card","border-card",""],default:""},closable:Boolean,addable:Boolean,modelValue:{type:[String,Number]},editable:Boolean,tabPosition:{type:String,values:["top","right","bottom","left"],default:"top"},beforeLeave:{type:re(Function),default:()=>!0},stretch:Boolean}),te=l=>Je(l)||Qe(l),Ra={[Ve]:l=>te(l),tabClick:(l,d)=>d instanceof Event,tabChange:l=>te(l),edit:(l,d)=>["remove","add"].includes(d),tabRemove:l=>te(l),tabAdd:()=>!0},Da=q({name:"ElTabs",props:Ba,emits:Ra,setup(l,{emit:d,slots:h,expose:v}){var r;const a=J("tabs"),p=G(()=>["left","right"].includes(l.tabPosition)),{children:b,addChild:f,removeChild:w}=ka(ce(),"ElTabPane"),_=$(),o=$((r=l.modelValue)!=null?r:"0"),m=async(s,k=!1)=>{var x,B,N;if(!(o.value===s||me(s)))try{await((x=l.beforeLeave)==null?void 0:x.call(l,s,o.value))!==!1&&(o.value=s,k&&(d(Ve,s),d("tabChange",s)),(N=(B=_.value)==null?void 0:B.removeFocus)==null||N.call(B))}catch{}},u=(s,k,x)=>{s.props.disabled||(m(k,!0),d("tabClick",s,x))},E=(s,k)=>{s.props.disabled||me(s.props.name)||(k.stopPropagation(),d("edit",s.props.name,"remove"),d("tabRemove",s.props.name))},V=()=>{d("edit",void 0,"add"),d("tabAdd")};I(()=>l.modelValue,s=>m(s)),I(o,async()=>{var s;await de(),(s=_.value)==null||s.scrollToActiveTab()}),He(X,{props:l,currentName:o,registerPane:s=>{b.value.push(s)},sortPane:f,unregisterPane:w}),v({currentName:o});const C=({render:s})=>s();return()=>{const s=h["add-icon"],k=l.editable||l.addable?t("div",{class:[a.e("new-tab"),p.value&&a.e("new-tab-vertical")],tabindex:"0",onClick:V,onKeydown:N=>{N.code===H.enter&&V()}},[s?oe(h,"add-icon"):t(Y,{class:a.is("icon-plus")},{default:()=>[t(We,null,null)]})]):null,x=t("div",{class:[a.e("header"),p.value&&a.e("header-vertical"),a.is(l.tabPosition)]},[t(C,{render:()=>{const N=b.value.some(R=>R.slots.label);return t(za,{ref:_,currentName:o.value,editable:l.editable,type:l.type,panes:b.value,stretch:l.stretch,onTabClick:u,onTabRemove:E},{$stable:!N})}},null),k]),B=t("div",{class:a.e("content")},[oe(h,"default")]);return t("div",{class:[a.b(),a.m(l.tabPosition),{[a.m("card")]:l.type==="card",[a.m("border-card")]:l.type==="border-card"}]},[B,x])}}}),Oa=W({label:{type:String,default:""},name:{type:[String,Number]},closable:Boolean,disabled:Boolean,lazy:Boolean}),ze="ElTabPane",Ua=q({name:ze}),Aa=q({...Ua,props:Oa,setup(l){const d=l,h=ce(),v=Xe(),r=ie(X);r||ue(ze,"usage: <el-tabs><el-tab-pane /></el-tabs/>");const a=J("tab-pane"),p=$(),b=G(()=>d.closable||r.props.closable),f=pe(()=>{var u;return r.currentName.value===((u=d.name)!=null?u:p.value)}),w=$(f.value),_=G(()=>{var u;return(u=d.name)!=null?u:p.value}),o=pe(()=>!d.lazy||w.value||f.value);I(f,u=>{u&&(w.value=!0)});const m=Ze({uid:h.uid,slots:v,props:d,paneName:_,active:f,index:p,isClosable:b});return r.registerPane(m),Ce(()=>{r.sortPane(m)}),ea(()=>{r.unregisterPane(m.uid)}),(u,E)=>T(o)?K((P(),L("div",{key:0,id:`pane-${T(_)}`,class:ye(T(a).b()),role:"tabpanel","aria-hidden":!T(f),"aria-labelledby":`tab-${T(_)}`},[oe(u.$slots,"default")],10,["id","aria-hidden","aria-labelledby"])),[[aa,T(f)]]):Ee("v-if",!0)}});var Be=he(Aa,[["__file","tab-pane.vue"]]);const Fa=ta(Da,{TabPane:Be}),Ma=la(Be);function ja(l){return Q({url:"/project",params:l})}function La(l){return Q({url:"/project",method:"post",data:l})}function Ia(l,d){return Q({url:"/project/"+l,method:"put",data:d})}function qa(l){return Q({url:"/project/"+l,method:"delete"})}const Ka={key:1},Ga={__name:"add-edit-dialog",props:{visible:{type:Boolean},visibleModifiers:{},formData:{type:Object},formDataModifiers:{}},emits:na(["done"],["update:visible","update:formData"]),setup(l,{emit:d}){const h=Te(),v=d,r=ve(l,"visible"),a=ve(l,"formData"),p=$(0),b=_=>{a.value.template.splice(_,1)},f=$(),w=async()=>{await f.value.validate();for(const _ in a.value.params)if(a.value.params[_].name==""||a.value.params[_].value=="")return ne.error("参数名和参数值必填");a.value.id==0?await La(a.value):await Ia(a.value.id,a.value),r.value=!1,v("done"),h.getProject(),ne.success("成功")};return(_,o)=>{const m=ke,u=we,E=ya,V=ha,C=Ne,s=xe,k=fa,x=va,B=Se,N=Ma,R=Fa,M=$e,c=ia;return P(),U(c,{title:a.value.id==0?"添加项目":"编辑项目",modelValue:r.value,"onUpdate:modelValue":o[12]||(o[12]=e=>r.value=e),"close-on-click-modal":!1,width:"50%"},{footer:i(()=>[t(C,{onClick:o[10]||(o[10]=e=>r.value=!1)},{default:i(()=>[A("取 消")]),_:1}),t(C,{type:"primary",onClick:o[11]||(o[11]=e=>w())},{default:i(()=>[A("确 定")]),_:1})]),default:i(()=>[t(M,{ref_key:"elFormRef",ref:f,model:a.value,"label-width":"160px"},{default:i(()=>[t(u,{label:"项目名",rules:{required:!0,trigger:"blur",message:"请输入项目名"},prop:"name"},{default:i(()=>[t(m,{modelValue:a.value.name,"onUpdate:modelValue":o[0]||(o[0]=e=>a.value.name=e)},null,8,["modelValue"])]),_:1}),t(u,{label:"项目描述",rules:{required:!0,trigger:"blur",message:"请输入项目描述"},prop:"desc"},{default:i(()=>[t(m,{modelValue:a.value.desc,"onUpdate:modelValue":o[1]||(o[1]=e=>a.value.desc=e)},null,8,["modelValue"])]),_:1}),t(u,{label:"GIT地址",rules:{required:!0,trigger:"blur",message:"请输入GIT地址"},prop:"git"},{default:i(()=>[t(m,{modelValue:a.value.git,"onUpdate:modelValue":o[2]||(o[2]=e=>a.value.git=e)},null,8,["modelValue"])]),_:1}),t(u,{label:"用户名",rules:{required:!0,trigger:"blur",message:"请输入用户名"},prop:"userName"},{default:i(()=>[t(m,{modelValue:a.value.userName,"onUpdate:modelValue":o[3]||(o[3]=e=>a.value.userName=e)},null,8,["modelValue"])]),_:1}),t(u,{label:"密码/token",rules:{required:a.value.id==0,trigger:"blur",message:"请输入密码"},prop:"token"},{default:i(()=>[t(m,{modelValue:a.value.token,"onUpdate:modelValue":o[4]||(o[4]=e=>a.value.token=e),placeholder:a.value.id==0?"":"不填写则不修改"},null,8,["modelValue","placeholder"])]),_:1},8,["rules"]),t(u,{label:"使用 TAG 更新版本",rules:{required:!0,trigger:"blur",type:"enum",enum:[0,1],message:"请设置使用TAG更新版本"},prop:"useTag"},{default:i(()=>[t(V,{modelValue:a.value.useTag,"onUpdate:modelValue":o[5]||(o[5]=e=>a.value.useTag=e)},{default:i(()=>[t(E,{label:0,border:""},{default:i(()=>[A("使用 latest")]),_:1}),t(E,{label:1,border:""},{default:i(()=>[A("使用指定 tag")]),_:1})]),_:1},8,["modelValue"])]),_:1}),t(u,{label:"预设变量"},{default:i(()=>[t(C,{type:"success",size:"small",onClick:o[6]||(o[6]=e=>a.value.params.push({}))},{default:i(()=>[A("新增")]),_:1}),t(B,{border:"",style:{"margin-top":"10px"},size:"small",data:a.value.params},{default:i(()=>[t(s,{align:"center",label:"参数名",prop:"name"},{default:i(({row:e})=>[t(m,{modelValue:e.name,"onUpdate:modelValue":n=>e.name=n,disabled:e.name=="tag",size:"small"},null,8,["modelValue","onUpdate:modelValue","disabled"])]),_:1}),t(s,{align:"center",label:"参数值",prop:"desc"},{default:i(({row:e})=>[e.name=="tag"?(P(),U(x,{key:0,size:"small",modelValue:e.value,"onUpdate:modelValue":n=>e.value=n,style:{width:"100%"}},{default:i(()=>[(P(!0),L(be,null,fe(e.options,(n,g)=>(P(),U(k,{key:g,label:n,value:n},null,8,["label","value"]))),128))]),_:2},1032,["modelValue","onUpdate:modelValue"])):(P(),U(m,{key:1,modelValue:e.value,"onUpdate:modelValue":n=>e.value=n,size:"small"},null,8,["modelValue","onUpdate:modelValue"]))]),_:1}),t(s,{align:"center",width:"170px",label:"操作"},{default:i(({$index:e})=>[t(C,{size:"small",type:"danger",onClick:n=>a.value.params.splice(e,1)},{default:i(()=>[A(" 删除 ")]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"])]),_:1}),t(u,{label:"项目k8s模板"},{default:i(()=>[t(C,{type:"success",size:"small",onClick:o[7]||(o[7]=e=>a.value.template.push({name:"新增模板",value:""}))},{default:i(()=>[A(" 新增 ")]),_:1}),a.value.template.length>0?(P(),U(R,{key:0,type:"border-card","tab-position":"left",modelValue:p.value,"onUpdate:modelValue":o[8]||(o[8]=e=>p.value=e),style:{"margin-top":"10px",width:"100%"},closable:"",onKeydownCapture:o[9]||(o[9]=sa(ra(()=>{},["stop"]),["delete"])),onTabRemove:b},{default:i(()=>[(P(!0),L(be,null,fe(a.value.template,(e,n)=>(P(),U(N,{key:n,label:e.name,name:n},{label:i(()=>[e.edit?(P(),U(m,{key:0,modelValue:e.name,"onUpdate:modelValue":g=>e.name=g,style:{width:"64px",padding:"0"},size:"small"},null,8,["modelValue","onUpdate:modelValue"])):(P(),L("div",Ka,ua(e.name),1)),t(C,{type:"primary",text:"",icon:e.edit?T(da):T(ca),onClick:g=>e.edit=!e.edit,size:"small"},null,8,["icon","onClick"])]),default:i(()=>[t(Ca,{modelValue:e.value,"onUpdate:modelValue":g=>e.value=g},null,8,["modelValue","onUpdate:modelValue"])]),_:2},1032,["label","name"]))),128))]),_:1},8,["modelValue"])):Ee("",!0)]),_:1})]),_:1},8,["model"])]),_:1},8,["title","modelValue"])}}},Ya=oa(Ga,[["__scopeId","data-v-a1523819"]]),Ha={class:"app-container"},Wa=se("p",{class:"page-title"},"项目列表",-1),Ja={class:"filter-container"},Qa={class:"page-container"},dt={__name:"index",setup(l){const d=Te(),h=()=>({id:0,name:"",desc:"",git:"",userName:"",token:"",useTag:"",params:[{name:"app_name",value:""},{name:"port",value:"8888"}],template:[]}),v=$(!1),r=$({}),a=async()=>{p.value=!0;const{data:C}=await ja(f.value);b.value=C.data,_.value=C.total,p.value=!1},{tableLoading:p,list:b,params:f,pageSizes:w,total:_,handleSizeChange:o,handleCurrentChange:m}=ga(a),u=()=>{v.value=!0,r.value=h()},E=async C=>{v.value=!0,Object.assign(r.value,C)},V=async C=>{await Ea.confirm("删除项目不可恢复","警告");const{message:s}=await qa(C.id);ne.success(s),a(),d.getProject()};return(C,s)=>{const k=Ne,x=_a,B=ke,N=we,R=$e,M=ba,c=xe,e=Se,n=pa,g=ma("permission"),z=Va;return P(),L("div",Ha,[Wa,se("div",Ja,[t(M,{gutter:20,justify:"space-between"},{default:i(()=>[t(x,{span:6},{default:i(()=>[K((P(),U(k,{type:"primary",size:"default",plain:"",onClick:s[0]||(s[0]=y=>u())},{default:i(()=>[A(" 添加项目 ")]),_:1})),[[g,"project:add"]])]),_:1}),t(x,{span:18},{default:i(()=>[t(R,{ref:"queryFormRef",inline:"",style:{float:"right"},size:"default",model:T(f)},{default:i(()=>[t(N,{label:"项目/描述",prop:"name"},{default:i(()=>[t(B,{modelValue:T(f).name,"onUpdate:modelValue":s[1]||(s[1]=y=>T(f).name=y),clearable:""},null,8,["modelValue"])]),_:1}),t(N,null,{default:i(()=>[t(k,{type:"primary",onClick:a},{default:i(()=>[A("查询")]),_:1})]),_:1})]),_:1},8,["model"])]),_:1})]),_:1})]),K((P(),U(e,{"element-loading-text":"加载中...",border:"",data:T(b)},{default:i(()=>[t(c,{label:"ID",prop:"id",align:"center",width:"60px"}),t(c,{align:"center",label:"项目名",prop:"name"}),t(c,{align:"center",label:"项目描述",prop:"desc"}),t(c,{align:"center",width:"170px",label:"操作"},{default:i(({row:y})=>[K((P(),U(k,{size:"small",type:"primary",onClick:O=>E(y)},{default:i(()=>[A("编辑 ")]),_:2},1032,["onClick"])),[[g,"project:edit"]]),K((P(),U(k,{size:"small",type:"danger",onClick:O=>V(y)},{default:i(()=>[A(" 删除 ")]),_:2},1032,["onClick"])),[[g,"project:del"]])]),_:1})]),_:1},8,["data"])),[[z,T(p)]]),se("div",Qa,[t(n,{background:"","current-page":T(f).page,"page-sizes":T(w),"page-size":T(f).pageSize,layout:"total, sizes, prev, pager, next, jumper",total:T(_),onSizeChange:T(o),onCurrentChange:T(m)},null,8,["current-page","page-sizes","page-size","total","onSizeChange","onCurrentChange"])]),t(Ya,{visible:v.value,"onUpdate:visible":s[2]||(s[2]=y=>v.value=y),formData:r.value,"onUpdate:formData":s[3]||(s[3]=y=>r.value=y),onDone:a},null,8,["visible","formData"])])}}};export{dt as default};
