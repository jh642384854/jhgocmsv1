(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-7a0f8fa0"],{"09f4":function(e,t,a){"use strict";a.d(t,"a",function(){return o}),Math.easeInOutQuad=function(e,t,a,i){return e/=i/2,e<1?a/2*e*e+t:(e--,-a/2*(e*(e-2)-1)+t)};var i=function(){return window.requestAnimationFrame||window.webkitRequestAnimationFrame||window.mozRequestAnimationFrame||function(e){window.setTimeout(e,1e3/60)}}();function r(e){document.documentElement.scrollTop=e,document.body.parentNode.scrollTop=e,document.body.scrollTop=e}function s(){return document.documentElement.scrollTop||document.body.parentNode.scrollTop||document.body.scrollTop}function o(e,t,a){var o=s(),l=e-o,n=20,c=0;t="undefined"===typeof t?500:t;var u=function e(){c+=n;var s=Math.easeInOutQuad(c,o,l,t);r(s),c<t?i(e):a&&"function"===typeof a&&a()};u()}},"49bc":function(e,t,a){"use strict";a.r(t);var i=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"app-container"},[a("el-form",{ref:"dataForm",attrs:{rules:e.rules,model:e.postForm,"label-position":"left","label-width":"100px"}},[a("select-material",{ref:"selectMaterial",on:{func:e.getMsgFormSon,changeMaterialType:e.doChangeMaterialType}}),e._v(" "),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:e.submitForm}},[e._v("保存")]),e._v(" "),a("el-button",[e._v("取消")])],1)],1)],1)},r=[],s=a("a78e"),o=a.n(s),l=a("7e70"),n=a("832d"),c={name:"wx",filters:{},components:{selectMaterial:n["a"]},beforeCreate:function(){this},data:function(){return{postForm:{wid:void 0,msgTypeID:"",materialID:""},rules:{msgTypeID:[{required:!0,message:"消息类型必须要选择",trigger:"blur"}],materialID:[{required:!0,message:"素材必须选择",trigger:"blur"}]}}},created:function(){this.getConfig()},methods:{getConfig:function(){var e=this,t=o.a.get(this.GLOBAL.WxCurrentIDToken),a={wid:t};Object(l["h"])(a).then(function(t){if(t.data.status==e.GLOBAL.SuccessText){e.postForm=t.data.results[0],e.$refs.selectMaterial.postForm.msgTypeID=t.data.results[0].msgTypeID,e.$refs.selectMaterial.postForm.materialID=t.data.results[0].materialID,e.$refs.selectMaterial.radioConfig.checkRowKey=t.data.results[0].materialID;var a=t.data.results[0].msgTypeID;switch(a){case"text":e.$refs.selectMaterial.selectMaterialContent=t.data.results[0].materialText;break;case"image":case"news":e.$refs.selectMaterial.selectMaterialTitle=t.data.results[0].materialText,e.$refs.selectMaterial.selectMaterialImage=t.data.results[0].materialPic;break;case"voice":e.$refs.selectMaterial.selectMaterialVoice=t.data.results[0].materialText;break;case"video":e.$refs.selectMaterial.selectMaterialVideo=t.data.results[0].materialText;break;default:e.$refs.selectMaterial.selectMaterialContent=t.data.results[0].materialText;break}}else e.GLOBAL.msgNotify("error","失败",t.data.msg)})},doChangeMaterialType:function(e){this.postForm.msgTypeID=e,this.postForm.materialID=""},getMsgFormSon:function(e){this.postForm.msgTypeID=e.msgTypeID,this.postForm.materialID=e.materialID},submitForm:function(){var e=this;this.$refs["dataForm"].validate(function(t){t&&(e.postForm["wid"]=o.a.get(e.GLOBAL.WxCurrentIDToken),Object(l["m"])(e.postForm).then(function(t){t.data.status==e.GLOBAL.SuccessText?e.GLOBAL.msgNotify("success","成功",t.data.msg):e.GLOBAL.msgNotify("error","失败",t.data.msg)}))})}}},u=c,m=a("2877"),d=Object(m["a"])(u,i,r,!1,null,null,null);t["default"]=d.exports},"7e70":function(e,t,a){"use strict";a.d(t,"d",function(){return r}),a.d(t,"a",function(){return s}),a.d(t,"n",function(){return o}),a.d(t,"c",function(){return l}),a.d(t,"f",function(){return n}),a.d(t,"e",function(){return c}),a.d(t,"m",function(){return u}),a.d(t,"h",function(){return m}),a.d(t,"l",function(){return d}),a.d(t,"g",function(){return f}),a.d(t,"i",function(){return p}),a.d(t,"b",function(){return g}),a.d(t,"k",function(){return h}),a.d(t,"j",function(){return v});var i=a("b775");function r(e){return Object(i["a"])({url:"admin/wx/list",method:"get",params:e})}function s(e){return Object(i["a"])({url:"admin/wx/creategzh",method:"post",data:e})}function o(e){return Object(i["a"])({url:"admin/wx/updategzh",method:"post",data:e})}function l(e){return Object(i["a"])({url:"admin/wx/deletegzh",method:"post",data:e})}function n(e){return Object(i["a"])({url:"admin/wx/getonegzh",method:"get",params:e})}function c(e){return Object(i["a"])({url:"admin/wx/getmaterial",method:"get",params:e})}function u(e){return Object(i["a"])({url:"admin/wx/configWelcomeMsg",method:"post",data:e})}function m(e){return Object(i["a"])({url:"admin/wx/configWelcomeMsg",method:"get",params:e})}function d(e){return Object(i["a"])({url:"admin/wx/configUnknowMsg",method:"post",data:e})}function f(e){return Object(i["a"])({url:"admin/wx/configUnknowMsg",method:"get",params:e})}function p(e){return Object(i["a"])({url:"admin/wx/listdiyreply",method:"get",params:e})}function g(e){return Object(i["a"])({url:"admin/wx/deletediyreply",method:"post",data:e})}function h(e){return Object(i["a"])({url:"admin/wx/savediyreply",method:"post",data:e})}function v(e){return Object(i["a"])({url:"admin/wx/getonediyreply",method:"get",params:e})}},"832d":function(e,t,a){"use strict";var i,r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[a("el-form-item",{attrs:{label:"消息类型",prop:"msgTypeID"}},[a("el-radio-group",{on:{change:e.clearSelectMaterial},model:{value:e.postForm.msgTypeID,callback:function(t){e.$set(e.postForm,"msgTypeID",t)},expression:"postForm.msgTypeID"}},e._l(e.msgType,function(t){return a("el-radio",{key:t.value,attrs:{label:t.value}},[e._v(e._s(t.name))])}),1)],1),e._v(" "),a("el-form-item",{attrs:{label:"选择素材",prop:"materialID"}},[""==e.postForm.materialID?a("div",{staticClass:"el-upload el-upload--picture-card",on:{click:e.selectMaterial}},[a("i",{staticClass:"el-icon-plus"})]):a("div",{staticClass:"selectMaterial"},[a("el-row",{attrs:{gutter:20}},[a("el-col",{attrs:{span:8}},["text"==e.postForm.msgTypeID?a("div",{staticClass:"materialShow"},[e._v(e._s(e.selectMaterialContent))]):"image"==e.postForm.msgTypeID?a("div",{staticClass:"materialShow"},[""!=e.selectMaterialImage?a("div",{staticClass:"block"},[a("el-image",{attrs:{src:e.selectMaterialImage}})],1):e._e()]):"news"==e.postForm.msgTypeID?a("div",{staticClass:"materialShow"},[""!=e.selectMaterialImage?a("div",{staticClass:"block"},[a("span",{staticClass:"demonstration"},[e._v(e._s(e.selectMaterialTitle))]),a("br"),e._v(" "),a("el-image",{attrs:{src:e.selectMaterialImage}})],1):e._e()]):"voice"==e.postForm.msgTypeID?a("div",{staticClass:"materialShow"},[""!=e.selectMaterialVoice?a("div",{staticClass:"block"},[a("span",{staticClass:"iconfont icon-MP"}),e._v(" "+e._s(e.selectMaterialVoice)+"\n            ")]):e._e()]):"video"==e.postForm.msgTypeID?a("div",{staticClass:"materialShow"},[""!=e.selectMaterialVideo?a("div",{staticClass:"block"},[a("span",{staticClass:"iconfont icon-video"}),e._v(" "+e._s(e.selectMaterialVideo)+"\n            ")]):e._e()]):e._e()]),e._v(" "),a("el-col",{attrs:{span:4}},[a("el-button",{attrs:{type:"success"},on:{click:e.selectMaterial}},[e._v("重新选择")])],1)],1)],1),e._v(" "),a("el-input",{attrs:{type:"hidden"},model:{value:e.postForm.materialID,callback:function(t){e.$set(e.postForm,"materialID",t)},expression:"postForm.materialID"}})],1),e._v(" "),a("el-dialog",{attrs:{title:e.materialTitle,visible:e.dialogTextVisible},on:{"update:visible":function(t){e.dialogTextVisible=t}}},[a("vxe-table",{ref:"xTable",staticClass:"radio-table",attrs:{border:"",height:"300",data:e.gridData,"row-id":"media_id","radio-config":e.radioConfig,"row-class-name":e.rowClassName},on:{"radio-change":e.radioChangeEvent}},[a("vxe-table-column",{attrs:{type:"radio",width:"60"},scopedSlots:e._u([{key:"header",fn:function(){return[a("vxe-button",{attrs:{type:"text",disabled:!e.selectTextRow},on:{click:e.clearRadioRowEevnt}},[e._v("取消")])]},proxy:!0}])}),e._v(" "),a("vxe-table-column",{attrs:{field:"id",width:"100",title:"编号"}}),e._v(" "),a("vxe-table-column",{attrs:{field:"content",title:"内容"}})],1),e._v(" "),a("pagination",{directives:[{name:"show",rawName:"v-show",value:e.total>0,expression:" total>0 "}],attrs:{total:e.total},on:{pagination:e.getMaterialList}}),e._v(" "),a("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:function(t){e.dialogTextVisible=!1}}},[e._v("\n        取消\n      ")]),e._v(" "),a("el-button",{attrs:{type:"primary"},on:{click:e.confirmSelectMaterial}},[e._v("\n        确定\n      ")])],1)],1),e._v(" "),a("el-dialog",{attrs:{title:e.materialTitle,visible:e.dialogImageVisible},on:{"update:visible":function(t){e.dialogImageVisible=t}}},[a("el-row",{staticClass:"mb10"},[a("el-radio-group",{on:{change:e.getImageMaterial},model:{value:e.selectMaterialRow,callback:function(t){e.selectMaterialRow=t},expression:"selectMaterialRow"}},e._l(e.gridData,function(t,i){return a("el-col",{key:t.id,staticClass:"mb10",staticStyle:{"margin-right":"20px"},attrs:{span:5}},[a("el-card",{attrs:{"body-style":{padding:"5px"},shadow:"hover"}},[a("div",{staticClass:"block",staticStyle:{"text-align":"center"}},[a("el-image",{staticStyle:{width:"180px",height:"180px"},attrs:{src:t.url,fit:"fill"}})],1),e._v(" "),a("div",{staticStyle:{padding:"14px"}},[a("span",[a("el-radio",{attrs:{label:t.media_id},model:{value:e.postForm.materialID,callback:function(t){e.$set(e.postForm,"materialID",t)},expression:"postForm.materialID"}},[e._v(e._s(t.name))])],1)])])],1)}),1)],1),e._v(" "),a("pagination",{directives:[{name:"show",rawName:"v-show",value:e.total>0,expression:" total>0 "}],attrs:{total:e.total},on:{pagination:e.getMaterialList}}),e._v(" "),a("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:function(t){e.dialogImageVisible=!1}}},[e._v("\n        取消\n      ")]),e._v(" "),a("el-button",{attrs:{type:"primary"},on:{click:e.confirmSelectMaterial}},[e._v("\n        确定\n      ")])],1)],1)],1)},s=[],o=(a("7f7f"),a("ac6a"),a("a78e"),a("c695"),a("7e70")),l=a("333d"),n={name:"wx",filters:{},components:{Pagination:l["a"]},beforeCreate:function(){i=this},data:function(){return{total:0,gridData:[],listQuery:{page:1,limit:20,type:""},radioConfig:{trigger:"row",checkRowKey:""},selectMaterialContent:"",selectMaterialTitle:"",selectMaterialImage:"",selectMaterialVoice:"",selectMaterialVideo:"",dialogTextVisible:!1,dialogImageVisible:!1,materialTitle:"选择素材",msgType:[{name:"文本",value:"text"},{name:"图片",value:"image"},{name:"图文",value:"news"},{name:"音频",value:"voice"},{name:"视频",value:"video"}],selectTextRow:null,selectMaterialRow:"",postForm:{msgTypeID:"",materialID:""}}},created:function(){},methods:{selectMaterial:function(){if(""==this.postForm.msgTypeID)this.$message.error("请先选择消息类型在选择素材");else switch(this.getMaterialList(),this.postForm.msgTypeID){case"text":case"voice":case"video":this.dialogTextVisible=!0;break;case"image":case"news":this.dialogImageVisible=!0;break;default:this.dialogTextVisible=!0;break}},getMaterialList:function(){var e=this;this.listQuery.type=this.postForm.msgTypeID,Object(o["e"])(this.listQuery).then(function(t){e.gridData=t.data.items,e.total=t.data.total})},radioChangeEvent:function(e){var t=e.row;this.selectTextRow=t},rowClassName:function(e){var t=e.row;return{"row-checked":this.selectTextRow===t}},clearRadioRowEevnt:function(){this.selectTextRow=null,this.$refs.xTable.clearRadioRow()},getImageMaterial:function(e){this.selectMaterialRow=e},clearSelectMaterial:function(){this.selectMaterialContent="",this.selectMaterialTitle="",this.selectMaterialImage="",this.selectMaterialVoice="",this.selectMaterialVideo="",this.$emit("changeMaterialType",this.postForm.msgTypeID)},confirmSelectMaterial:function(){var e=this.postForm.msgTypeID;switch(e){case"text":case"voice":case"video":null!=this.selectTextRow?(this.postForm.materialID=this.selectTextRow.id,"text"==e?this.selectMaterialContent=this.selectTextRow.content:"voice"==e?this.selectMaterialVoice=this.selectTextRow.content:"video"==e&&(this.selectMaterialVideo=this.selectTextRow.content),this.dialogTextVisible=!1):this.$message.error("请先选择好素材");break;case"image":case"news":""!=this.selectMaterialRow?(this.gridData.forEach(function(e,t,a){e.media_id==i.selectMaterialRow&&(i.postForm.materialID=e.media_id,i.selectMaterialTitle=e.name,i.selectMaterialImage=e.url)}),this.dialogImageVisible=!1):this.$message.error("请先选择好素材");break;default:this.selectMaterialContent=this.selectTextRow.content,this.dialogTextVisible=!1;break}this.$emit("func",this.postForm)}}},c=n,u=a("2877"),m=Object(u["a"])(c,r,s,!1,null,null,null);t["a"]=m.exports},aa77:function(e,t,a){var i=a("5ca1"),r=a("be13"),s=a("79e5"),o=a("fdef"),l="["+o+"]",n="​",c=RegExp("^"+l+l+"*"),u=RegExp(l+l+"*$"),m=function(e,t,a){var r={},l=s(function(){return!!o[e]()||n[e]()!=n}),c=r[e]=l?t(d):o[e];a&&(r[a]=c),i(i.P+i.F*l,"String",r)},d=m.trim=function(e,t){return e=String(r(e)),1&t&&(e=e.replace(c,"")),2&t&&(e=e.replace(u,"")),e};e.exports=m},c5f6:function(e,t,a){"use strict";var i=a("7726"),r=a("69a8"),s=a("2d95"),o=a("5dbc"),l=a("6a99"),n=a("79e5"),c=a("9093").f,u=a("11e9").f,m=a("86cc").f,d=a("aa77").trim,f="Number",p=i[f],g=p,h=p.prototype,v=s(a("2aeb")(h))==f,b="trim"in String.prototype,w=function(e){var t=l(e,!1);if("string"==typeof t&&t.length>2){t=b?t.trim():d(t,3);var a,i,r,s=t.charCodeAt(0);if(43===s||45===s){if(a=t.charCodeAt(2),88===a||120===a)return NaN}else if(48===s){switch(t.charCodeAt(1)){case 66:case 98:i=2,r=49;break;case 79:case 111:i=8,r=55;break;default:return+t}for(var o,n=t.slice(2),c=0,u=n.length;c<u;c++)if(o=n.charCodeAt(c),o<48||o>r)return NaN;return parseInt(n,i)}}return+t};if(!p(" 0o1")||!p("0b1")||p("+0x1")){p=function(e){var t=arguments.length<1?0:e,a=this;return a instanceof p&&(v?n(function(){h.valueOf.call(a)}):s(a)!=f)?o(new g(w(t)),a,p):w(t)};for(var M,I=a("9e1e")?c(g):"MAX_VALUE,MIN_VALUE,NaN,NEGATIVE_INFINITY,POSITIVE_INFINITY,EPSILON,isFinite,isInteger,isNaN,isSafeInteger,MAX_SAFE_INTEGER,MIN_SAFE_INTEGER,parseFloat,parseInt,isInteger".split(","),T=0;I.length>T;T++)r(g,M=I[T])&&!r(p,M)&&m(p,M,u(g,M));p.prototype=h,h.constructor=p,a("2aba")(i,f,p)}},fdef:function(e,t){e.exports="\t\n\v\f\r   ᠎             　\u2028\u2029\ufeff"}}]);