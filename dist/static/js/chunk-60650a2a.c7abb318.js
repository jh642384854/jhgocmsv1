(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-60650a2a"],{"0171":function(t,e,a){},"09f4":function(t,e,a){"use strict";a.d(e,"a",function(){return n}),Math.easeInOutQuad=function(t,e,a,i){return t/=i/2,t<1?a/2*t*t+e:(t--,-a/2*(t*(t-2)-1)+e)};var i=function(){return window.requestAnimationFrame||window.webkitRequestAnimationFrame||window.mozRequestAnimationFrame||function(t){window.setTimeout(t,1e3/60)}}();function r(t){document.documentElement.scrollTop=t,document.body.parentNode.scrollTop=t,document.body.scrollTop=t}function o(){return document.documentElement.scrollTop||document.body.parentNode.scrollTop||document.body.scrollTop}function n(t,e,a){var n=o(),s=t-n,l=20,u=0;e="undefined"===typeof e?500:e;var c=function t(){u+=l;var o=Math.easeInOutQuad(u,n,s,e);r(o),u<e?i(t):a&&"function"===typeof a&&a()};c()}},"456d":function(t,e,a){var i=a("4bf8"),r=a("0d58");a("5eda")("keys",function(){return function(t){return r(i(t))}})},"4cce":function(t,e,a){"use strict";a.r(e);var i,r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"app-container"},[a("div",{staticClass:"filter-container"},[a("el-button",{staticClass:"filter-item mr10",attrs:{type:"primary",icon:"el-icon-edit"},on:{click:function(e){return t.handleCreate()}}},[t._v("\n        创建\n      ")]),t._v(" "),a("el-button",{staticClass:"filter-item mr10",attrs:{type:"danger",icon:"el-icon-delete"},on:{click:function(e){return t.batchDelete()}}},[t._v("\n        删除\n      ")])],1),t._v(" "),a("vxe-table",{ref:"xTable",attrs:{size:t.GLOBAL.VxeTableSize,data:t.tableData,loading:t.listLoading},on:{"update:data":function(e){t.tableData=e}}},[a("vxe-table-column",{attrs:{type:"selection",width:"60"}}),t._v(" "),a("vxe-table-column",{attrs:{type:"index",width:"60"}}),t._v(" "),a("vxe-table-column",{attrs:{field:"name",title:"广告名称"}}),t._v(" "),a("vxe-table-column",{attrs:{field:"type",title:"广告类别"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v("\n          "+t._s(t._f("advCateFilter")(e.row.cid))+"\n        ")]}}])}),t._v(" "),a("vxe-table-column",{attrs:{field:"list_order",title:"排序",width:"100"}}),t._v(" "),a("vxe-table-column",{attrs:{field:"img_url",title:"图片地址",width:"120"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-image",{staticStyle:{width:"100px",height:"100px"},attrs:{src:e.row.img_url,fit:"contain"},on:{click:function(a){return t.popImg(e.row)}}})]}}])}),t._v(" "),a("vxe-table-column",{attrs:{field:"link_url",title:"链接地址","show-overflow":""},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-link",{attrs:{href:e.row.link_url,target:"_blank",type:"primary"}},[t._v(t._s(e.row.link_url))])]}}])}),t._v(" "),a("vxe-table-column",{attrs:{field:"created_at",title:"创建时间"}}),t._v(" "),a("vxe-table-column",{attrs:{title:"操作"},scopedSlots:t._u([{key:"default",fn:function(e){var i=e.row;return[a("el-button",{attrs:{size:"mini",type:"success"}},[a("span",{on:{click:function(e){return t.handleEdit(i.id,i)}}},[t._v("编辑")])]),t._v(" "),a("el-button",{attrs:{size:"mini",type:"danger"},on:{click:function(e){return t.handleDelete(i)}}},[t._v("删除")])]}}])})],1),t._v(" "),a("pagination",{directives:[{name:"show",rawName:"v-show",value:t.total>0,expression:" total>0 "}],attrs:{total:t.total},on:{pagination:t.getList}}),t._v(" "),a("el-dialog",{attrs:{title:t.textMap[t.dialogStatus],visible:t.dialogFormVisible},on:{"update:visible":function(e){t.dialogFormVisible=e}}},[a("el-form",{ref:"dataForm",attrs:{rules:t.rules,model:t.postForm,"label-position":"left","label-width":"100px"}},[a("el-form-item",{attrs:{label:"广告标题",prop:"name"}},[a("el-input",{model:{value:t.postForm.name,callback:function(e){t.$set(t.postForm,"name",e)},expression:"postForm.name"}})],1),t._v(" "),a("el-form-item",{attrs:{label:"广告类别",prop:"cid"}},[a("el-select",{attrs:{placeholder:"请选择"},model:{value:t.postForm.cid,callback:function(e){t.$set(t.postForm,"cid",e)},expression:"postForm.cid"}},t._l(t.advcategories,function(t){return a("el-option",{key:t.id,attrs:{label:t.name,value:t.id}})}),1)],1),t._v(" "),a("el-row",{attrs:{gutter:20}},[a("el-col",{attrs:{span:8}},[a("el-form-item",{attrs:{label:"排序",prop:"list_order"}},[a("el-input-number",{attrs:{min:1,max:100},model:{value:t.postForm.list_order,callback:function(e){t.$set(t.postForm,"list_order",e)},expression:"postForm.list_order"}})],1)],1),t._v(" "),a("el-col",{attrs:{span:16}},[a("el-form-item",{attrs:{label:"链接地址",prop:"link_url"}},[a("el-input",{model:{value:t.postForm.link_url,callback:function(e){t.$set(t.postForm,"link_url",t._n(e))},expression:"postForm.link_url"}})],1)],1)],1),t._v(" "),a("el-form-item",{attrs:{label:"广告图片",prop:"img_url"}},[a("el-upload",{ref:"upload",staticClass:"thumb-uploader",attrs:{accept:t.GLOBAL.AllowImageSuffix,action:t.GLOBAL.UploadAttachmentUrl,"list-type":"picture","auto-upload":!1,"show-file-list":!1,"on-success":t.handleThumbSuccess,"on-change":t.imgPreview}},[t.postForm.img_url?a("img",{staticClass:"thumb",attrs:{src:t.postForm.img_url}}):a("i",{staticClass:"el-icon-plus thumb-uploader-icon"}),t._v(" "),a("div",{staticClass:"el-upload__tip",attrs:{slot:"tip"},slot:"tip"},[t._v("只能上传jpg/png文件，且不超过500kb")])])],1)],1),t._v(" "),a("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:function(e){t.dialogFormVisible=!1}}},[t._v("\n          取消\n        ")]),t._v(" "),a("el-button",{attrs:{type:"primary"},on:{click:function(e){"create"===t.dialogStatus?t.createData():t.updateData()}}},[t._v("\n          "+t._s(t.textMap[t.dialogStatus])+"\n        ")])],1)],1),t._v(" "),a("el-dialog",{staticStyle:{"text-align":"center"},attrs:{visible:t.dialogImgVisible},on:{"update:visible":function(e){t.dialogImgVisible=e}}},[a("img",{staticClass:"dialogImg",attrs:{src:t.dialogImageUrl,alt:""}})])],1)},o=[],n=(a("ac4d"),a("8a81"),a("ac6a"),a("456d"),a("7f7f"),a("c695"),a("fe1d")),s=a("333d"),l={name:"advs",components:{Pagination:s["a"]},filters:{advCateFilter:function(t){var e="";for(var a in i.advcategories)if(i.advcategories[a].id==t){e=i.advcategories[a].name;break}return e}},beforeCreate:function(){i=this},data:function(){return{dialogImageUrl:"",dialogImgVisible:!1,hasSelectFileUpload:!1,listLoading:!0,total:0,tableData:[],advcategories:[],listQuery:{page:1,limit:20,cid:""},sorts:[],dialogStatus:"",dialogFormVisible:!1,textMap:{update:"修改",create:"创建"},postForm:{name:"",cid:"",img_url:"",upload_img_url:"",list_order:"",link_url:"",created_at:""},rules:{name:[{required:!0,message:"广告名称必须填写",trigger:"blur"}],cid:[{required:!0,message:"广告类别必须选择",trigger:"blur"}],link_url:[{required:!0,message:"广告链接必须填写",trigger:"blur"}],img_url:[{required:!0,message:"广告图片必须上传",trigger:"blur"}]}}},created:function(){this.handerQuery(),this.getList(),this.getAdvCategories()},methods:{handerQuery:function(){Object.keys(this.$route.query).length>0&&"undefined"!=typeof this.$route.query.cid&&parseInt(this.$route.query.cid)>0&&(this.listQuery.cid=this.$route.query.cid)},getList:function(){var t=this;this.listLoading=!0,Object(n["g"])(this.listQuery).then(function(e){t.tableData=e.data.items,t.total=e.data.total,t.listLoading=!1})},getAdvCategories:function(){var t=this;Object(n["f"])().then(function(e){t.advcategories=e.data.items})},formatterImgurl:function(t){t.cellValue;var e=t.row;t.rowIndex,t.column,t.columnIndex;return e.width+"*"+e.height+"px"},popImg:function(t){""!=t.img_url?(this.dialogImgVisible=!0,this.dialogImageUrl=t.img_url):this.GLOBAL.msgNotify("error","失败","没有上传图片，无法预览")},imgPreview:function(t,e){this.GLOBAL.uploadImageCheck(t,e)&&(this.hasSelectFileUpload=!0,this.postForm.img_url=t.url)},handleThumbSuccess:function(t,e,a){if(t.data.status!=this.GLOBAL.SuccessText)return this.GLOBAL.msgNotify("error","失败",t.data.msg),!1;this.postForm.upload_img_url=t.data.attachment.upload_file_path},resetTemp:function(){this.postForm={name:"",cid:"",img_url:"",upload_img_url:"",list_order:"",link_url:"",created_at:""}},handleCreate:function(){var t=this;this.resetTemp(),this.dialogStatus="create",this.dialogFormVisible=!0,this.$nextTick(function(){t.$refs["dataForm"].clearValidate()})},createData:function(){var t=this;this.$refs["dataForm"].validate(function(e){e&&(t.hasSelectFileUpload&&t.$refs.upload.submit(),Object(n["b"])(t.postForm).then(function(e){e.data.status==t.GLOBAL.SuccessText?(""!=t.postForm.upload_img_url&&(t.postForm.img_url=t.postForm.upload_img_url),t.postForm.id=e.data.last_id,t.postForm.created_at=e.data.created_at,t.tableData.unshift(t.postForm),t.dialogFormVisible=!1,t.GLOBAL.msgNotify("success","成功",e.data.msg)):t.GLOBAL.msgNotify("error","失败",e.data.msg)}))})},handleEdit:function(t,e){var a=this;this.postForm=Object.assign({},e),this.dialogStatus="update",this.dialogFormVisible=!0,this.$nextTick(function(){a.$refs["dataForm"].clearValidate()})},updateData:function(){var t=this;this.$refs["dataForm"].validate(function(e){if(e){t.hasSelectFileUpload&&t.$refs.upload.submit();var a=Object.assign({},t.postForm);Object(n["i"])(a).then(function(e){if(e.data.status==t.GLOBAL.SuccessText){"undefined"!=typeof t.postForm.upload_img_url&&(t.postForm.img_url=t.postForm.upload_img_url);var a=!0,i=!1,r=void 0;try{for(var o,n=t.tableData[Symbol.iterator]();!(a=(o=n.next()).done);a=!0){var s=o.value;if(s.id===t.postForm.id){var l=t.tableData.indexOf(s);t.tableData.splice(l,1,t.postForm);break}}}catch(u){i=!0,r=u}finally{try{a||null==n.return||n.return()}finally{if(i)throw r}}t.dialogFormVisible=!1,t.GLOBAL.msgNotify("success","成功",e.data.msg)}else t.GLOBAL.msgNotify("error","失败",e.data.msg)})}})},batchDelete:function(){var t=this,e=this.$refs.xTable.getSelectRecords();if(e.length){var a=[];for(var i in e)a.push(e[i].id);var r={ids:a};Object(n["d"])(r).then(function(e){e.data.status==t.GLOBAL.SuccessText?(t.GLOBAL.msgNotify("success","成功",e.data.msg),t.$refs.xTable.removeSelecteds()):t.GLOBAL.msgNotify("error","失败",e.data.msg)})}else this.$alert("请至少选择一条数据！")},handleDelete:function(t){var e=this;this.$confirm("确定要删除该记录吗？","确认信息",{distinguishCancelAndClose:!0,confirmButtonText:"删除",cancelButtonText:"放弃"}).then(function(){var a={ids:t.id};Object(n["d"])(a).then(function(a){if(a.data.status==e.GLOBAL.SuccessText){e.GLOBAL.msgNotify("success","成功",a.data.msg);var i=e.tableData.indexOf(t);e.tableData.splice(i,1)}else e.GLOBAL.msgNotify("error","失败",a.data.msg)})}).catch(function(t){e.$message({type:"info",message:"cancel"===t?"放弃操作并离开页面":"停留在当前页面"})})}}},u=l,c=(a("f9ba"),a("2877")),d=Object(c["a"])(u,r,o,!1,null,"4bdefbf2",null);e["default"]=d.exports},"5eda":function(t,e,a){var i=a("5ca1"),r=a("8378"),o=a("79e5");t.exports=function(t,e){var a=(r.Object||{})[t]||Object[t],n={};n[t]=e(a),i(i.S+i.F*o(function(){a(1)}),"Object",n)}},aa77:function(t,e,a){var i=a("5ca1"),r=a("be13"),o=a("79e5"),n=a("fdef"),s="["+n+"]",l="​",u=RegExp("^"+s+s+"*"),c=RegExp(s+s+"*$"),d=function(t,e,a){var r={},s=o(function(){return!!n[t]()||l[t]()!=l}),u=r[t]=s?e(m):n[t];a&&(r[a]=u),i(i.P+i.F*s,"String",r)},m=d.trim=function(t,e){return t=String(r(t)),1&e&&(t=t.replace(u,"")),2&e&&(t=t.replace(c,"")),t};t.exports=d},c5f6:function(t,e,a){"use strict";var i=a("7726"),r=a("69a8"),o=a("2d95"),n=a("5dbc"),s=a("6a99"),l=a("79e5"),u=a("9093").f,c=a("11e9").f,d=a("86cc").f,m=a("aa77").trim,f="Number",p=i[f],g=p,v=p.prototype,h=o(a("2aeb")(v))==f,b="trim"in String.prototype,_=function(t){var e=s(t,!1);if("string"==typeof e&&e.length>2){e=b?e.trim():m(e,3);var a,i,r,o=e.charCodeAt(0);if(43===o||45===o){if(a=e.charCodeAt(2),88===a||120===a)return NaN}else if(48===o){switch(e.charCodeAt(1)){case 66:case 98:i=2,r=49;break;case 79:case 111:i=8,r=55;break;default:return+e}for(var n,l=e.slice(2),u=0,c=l.length;u<c;u++)if(n=l.charCodeAt(u),n<48||n>r)return NaN;return parseInt(l,i)}}return+e};if(!p(" 0o1")||!p("0b1")||p("+0x1")){p=function(t){var e=arguments.length<1?0:t,a=this;return a instanceof p&&(h?l(function(){v.valueOf.call(a)}):o(a)!=f)?n(new g(_(e)),a,p):_(e)};for(var F,y=a("9e1e")?u(g):"MAX_VALUE,MIN_VALUE,NaN,NEGATIVE_INFINITY,POSITIVE_INFINITY,EPSILON,isFinite,isInteger,isNaN,isSafeInteger,MAX_SAFE_INTEGER,MIN_SAFE_INTEGER,parseFloat,parseInt,isInteger".split(","),x=0;y.length>x;x++)r(g,F=y[x])&&!r(p,F)&&d(p,F,c(g,F));p.prototype=v,v.constructor=p,a("2aba")(i,f,p)}},f9ba:function(t,e,a){"use strict";var i=a("0171"),r=a.n(i);r.a},fdef:function(t,e){t.exports="\t\n\v\f\r   ᠎             　\u2028\u2029\ufeff"},fe1d:function(t,e,a){"use strict";a.d(e,"f",function(){return r}),a.d(e,"e",function(){return o}),a.d(e,"c",function(){return n}),a.d(e,"a",function(){return s}),a.d(e,"h",function(){return l}),a.d(e,"g",function(){return u}),a.d(e,"d",function(){return c}),a.d(e,"b",function(){return d}),a.d(e,"i",function(){return m});var i=a("b775");function r(t){return Object(i["a"])({url:"admin/advs/categories",method:"get",params:t})}function o(t){return Object(i["a"])({url:"admin/advs/catetypes",method:"get",params:t})}function n(t){return Object(i["a"])({url:"admin/advs/deletecate",method:"get",params:{id:t}})}function s(t){return Object(i["a"])({url:"admin/advs/createcate",method:"post",data:t})}function l(t){return Object(i["a"])({url:"admin/advs/updatecate",method:"post",data:t})}function u(t){return Object(i["a"])({url:"admin/advs/list",method:"get",params:t})}function c(t){return Object(i["a"])({url:"admin/advs/delete",method:"post",data:t})}function d(t){return Object(i["a"])({url:"admin/advs/create",method:"post",data:t})}function m(t){return Object(i["a"])({url:"admin/advs/update",method:"post",data:t})}}}]);