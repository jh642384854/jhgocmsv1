(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-d11d6d92"],{"5fd2":function(t,e,r){"use strict";var s=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"app-container"},[r("el-form",{ref:"dataForm",attrs:{rules:t.rules,model:t.postForm,"label-width":"80px"}},[r("el-form-item",{attrs:{label:"用户名",prop:"username"}},[r("el-input",{model:{value:t.postForm.username,callback:function(e){t.$set(t.postForm,"username",e)},expression:"postForm.username"}})],1),t._v(" "),r("el-form-item",{attrs:{label:"昵称",prop:"nickname"}},[r("el-input",{model:{value:t.postForm.nickname,callback:function(e){t.$set(t.postForm,"nickname",e)},expression:"postForm.nickname"}})],1),t._v(" "),r("el-form-item",{attrs:{label:"用户头像",prop:"avator"}},[r("el-upload",{ref:"upload",staticClass:"thumb-uploader",attrs:{accept:t.GLOBAL.AllowImageSuffix,action:t.GLOBAL.UploadAttachmentUrl,"list-type":"picture","auto-upload":!1,"show-file-list":!1,"on-success":t.handleThumbSuccess,"on-change":t.imgPreview}},[t.postForm.avator?r("img",{staticClass:"thumb",attrs:{src:t.postForm.avator}}):r("i",{staticClass:"el-icon-plus thumb-uploader-icon"}),t._v(" "),r("div",{staticClass:"el-upload__tip",attrs:{slot:"tip"},slot:"tip"},[t._v(t._s(t.GLOBAL.ImageUploadTips))])])],1),t._v(" "),r("el-form-item",{attrs:{label:"性别",prop:"sex"}},[r("el-radio-group",{model:{value:t.postForm.sex,callback:function(e){t.$set(t.postForm,"sex",e)},expression:"postForm.sex"}},t._l(t.userSex,function(e){return r("el-radio",{key:e.label,attrs:{label:e.label}},[t._v(t._s(e.value))])}),1)],1),t._v(" "),r("el-form-item",{attrs:{label:"用户分组",prop:"groupid"}},[r("el-radio-group",{model:{value:t.postForm.groupid,callback:function(e){t.$set(t.postForm,"groupid",e)},expression:"postForm.groupid"}},t._l(t.userGroups,function(e){return r("el-radio",{key:e.label,attrs:{label:e.label}},[t._v(t._s(e.value))])}),1)],1),t._v(" "),r("el-form-item",{attrs:{label:"用户来源",prop:"from"}},[r("el-radio-group",{model:{value:t.postForm.from,callback:function(e){t.$set(t.postForm,"from",e)},expression:"postForm.from"}},t._l(t.userFrom,function(e){return r("el-radio",{key:e.label,attrs:{label:e.label}},[t._v(t._s(e.value))])}),1)],1),t._v(" "),r("el-form-item",{attrs:{label:"用户状态",prop:"status"}},[r("el-radio-group",{model:{value:t.postForm.status,callback:function(e){t.$set(t.postForm,"status",e)},expression:"postForm.status"}},t._l(t.userStatus,function(e){return r("el-radio",{key:e.label,attrs:{label:e.label}},[t._v(t._s(e.value))])}),1)],1),t._v(" "),r("el-form-item",{attrs:{label:"用户积分"}},[r("el-input-number",{attrs:{min:0,max:1e4},model:{value:t.postForm.points,callback:function(e){t.$set(t.postForm,"points",e)},expression:"postForm.points"}})],1),t._v(" "),r("el-form-item",[r("el-button",{attrs:{type:"primary"},on:{click:t.onSubmit}},[t._v("立即创建")]),t._v(" "),r("el-button",[t._v("取消")])],1)],1)],1)},a=[],o=(r("7f7f"),r("c24f")),i={id:void 0,username:"",nickname:"",avator:"",sex:"",groupid:"",from:"",points:"",status:""},u={name:"userForm",props:{isEdit:{type:Boolean,default:!1}},data:function(){return{hasSelectFileUpload:!1,newUploadFilePath:"",userStatus:[],userGroups:[],userSex:[],userFrom:[],postForm:Object.assign({},i),rules:{username:[{required:!0,message:"用户名必须填写",trigger:"blur"}],sex:[{required:!0,message:"请选择用户性别",trigger:"blur"}],groupid:[{required:!0,message:"请选择用户分组",trigger:"blur"}],status:[{required:!0,message:"请选择用户状态",trigger:"blur"}],from:[{required:!0,message:"请选择用户来源",trigger:"change"}]}}},created:function(){this.init()},methods:{init:function(){var t=this;if(this.isEdit){var e=this.$route.params&&this.$route.params.id;this.fetchUserData(e)}this.postForm.created_at=Date.parse(new Date),Object(o["d"])().then(function(e){t.userStatus=e.data.attrStatus,t.userGroups=e.data.attrGroups,t.userSex=e.data.attrSex,t.userFrom=e.data.attrFrom})},fetchUserData:function(t){var e=this;Object(o["c"])({id:t}).then(function(t){t.data.status==e.GLOBAL.SuccessText?e.postForm=t.data.results:e.$message.error(t.data.msg)})},onSubmit:function(){var t=this;this.$refs["dataForm"].validate(function(e){e&&(t.hasSelectFileUpload?t.$refs.upload.submit():t.saveData(t.postForm))})},saveData:function(t){var e=this;Object(o["e"])(this.postForm).then(function(t){t.data.status==e.GLOBAL.SuccessText?(e.GLOBAL.msgNotify("success","成功",t.data.msg),e.$router.push({path:"/user/list"})):e.$message.error(t.data.msg)})},handleThumbSuccess:function(t,e,r){t.data.status==this.GLOBAL.SuccessText?(this.postForm.upload_img_url=t.data.attachment.upload_file_path,console.log(this.postForm),this.saveData(this.postForm)):this.GLOBAL.msgNotify("error","失败",t.data.msg)},imgCheck:function(t,e){var r=t.name,s=/(.jpg|.jpeg|.gif|.png|.bmp)$/;return s.test(r.toLowerCase())?!(t.size>this.fileSizeLimit)||(this.$message.error("选择的图片大小不能超过500KB"),!1):(this.$message.error("请选择图片文件"),!1)},imgPreview:function(t,e){this.imgCheck(t,e)&&(this.hasSelectFileUpload=!0,this.postForm.avator=t.url)}}},l=u,n=r("2877"),m=Object(n["a"])(l,s,a,!1,null,null,null);e["a"]=m.exports},c24f:function(t,e,r){"use strict";r.d(e,"b",function(){return a}),r.d(e,"c",function(){return o}),r.d(e,"d",function(){return i}),r.d(e,"e",function(){return u}),r.d(e,"a",function(){return l});var s=r("b775");function a(t){return Object(s["a"])({url:"admin/user/list",method:"get",params:t})}function o(t){return Object(s["a"])({url:"admin/user/getone",method:"get",params:t})}function i(t){return Object(s["a"])({url:"admin/user/config",method:"get",params:t})}function u(t){return Object(s["a"])({url:"admin/user/save",method:"post",data:t})}function l(t){return Object(s["a"])({url:"admin/user/delete",method:"get",params:t})}},d00c:function(t,e,r){"use strict";r.r(e);var s=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("user-form",{attrs:{"is-edit":!0}})},a=[],o=r("5fd2"),i={name:"CreateUser",components:{UserForm:o["a"]}},u=i,l=r("2877"),n=Object(l["a"])(u,s,a,!1,null,null,null);e["default"]=n.exports}}]);