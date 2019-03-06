$(document).ready(function () {
    setInterval(hearthandler,50000)
    setInterval(resulthandler,1000)

    for (var i=1;i<=7;i++){
        if(i == 1){
            $('#detail-panel-'+i).show()
        }else{
            $('#detail-panel-'+i).hide()
        }
    }


    $('#left-nav-panel a:first').tab("show")
    $('#left-nav-panel').on('shown.bs.tab',function (e) {
        var target = $(e.target).find("a").attr("href");
        var id = target.substr(-1,1)
        console.log(id);
        for (var i=1;i<=7;i++){
            if(id == i){
                $('#detail-panel-'+i).show()
            }else{
                $('#detail-panel-'+i).hide()
            }
        }
    })


    GetUserInfo();
    LoadPlayerTable();
    LoadUserTable();
    LoadItemTable();
    LoadEquipTable();
    LoadNoticeTable();
    LoadChannelTable();
})


function GetUserInfo(){
    $.get("/user/info",{token:GetToken()},function(res){
        if (res!=null && res.code ==0){
            $("#user-info-name").text("欢迎登陆:"+res.data.username)
            $("#user-info-permission").text("当前权限:" +res.data.permission)
        }
    });
}


function LoadUserTable(){
    $("#user-table").bootstrapTable("destroy");
    $("#user-table").bootstrapTable({
        url:"/user/alluserinfo",
        method:"get",
        dataType:"json",
        pagination:true,
        clickToSelect: true,                //是否启用点击选中行
        sidePagination: "client",           //分页方式：client客户端分页，server服务端分页（*）
        pageNumber:1,                       //初始化加载第一页，默认第一页
        pageSize: 10,                       //每页的记录行数（*）
        pageList: [10, 25, 50, 100],        //可供选择的每页的行数（*）
        showRefresh: true,                  //是否显示刷新按钮
        search: true,                      //显示搜索框
        locale:"zh-CN",
        queryParams:{token:GetToken()},
        columns: [
            { field: 'id', title: 'ID',align:'center' },
            { field: 'username', title: '账号' ,align:'center' },
            { field: 'permission', title: '权限' ,align:'center' },
            { field: 'created', title: '创建于',align:'center'  },
            { field: 'id', title: '操作' ,formatter:'<a href="#">修改权限 </a><a href="#">修改密码</a>',align:'center' },
        ],
        responseHandler:function(res){
            if(res.code == 0){
                var temp =[]
                var d = res.data
                for(var i=0;i<d.length;i++){
                    temp.push({"id":d[i].id,"username":d[i].username,"permission":d[i].permission,"created":d[i].created})
                }
                return temp
            }
            return [] 
        },
    });
}

function LoadItemTable(){
    $("#item-table").bootstrapTable("destroy");
    $("#item-table").bootstrapTable({
        url:"/goods/alliteminfo",
        method:"get",
        dataType:"json",
        pagination:true,
        clickToSelect: true,                //是否启用点击选中行
        sidePagination: "client",           //分页方式：client客户端分页，server服务端分页（*）
        pageNumber:1,                       //初始化加载第一页，默认第一页
        pageSize: 20,                       //每页的记录行数（*）
        pageList: [10, 25, 50, 100],        //可供选择的每页的行数（*）
        showRefresh: true,                  //是否显示刷新按钮
        search: true,                      //显示搜索框
        locale:"zh-CN",
        queryParams:{token:GetToken()},
        columns: [
            { field: 'id', title: 'ID',align:'center' },
            { field: 'itemname', title: '名称' ,align:'center' },
            { field: 'desc', title: '描述' ,align:'center' },
        ],
        responseHandler:function(res){
            if(res.code == 0){
                var temp =[]
                var data = res.data
                for(var i=0;i<data.length;i++){
                    temp.push({"id":data[i].id,"itemname":data[i].name,"desc":data[i].desc})
                }
                return temp
            }
            return []
        },
    });
}

function LoadEquipTable(){
    $("#equip-table").bootstrapTable("destroy");
    $("#equip-table").bootstrapTable({
        url:"/goods/allequipinfo",
        method:"get",
        dataType:"json",
        pagination:true,
        clickToSelect: true,                //是否启用点击选中行
        sidePagination: "client",           //分页方式：client客户端分页，server服务端分页（*）
        pageNumber:1,                       //初始化加载第一页，默认第一页
        pageSize: 20,                       //每页的记录行数（*）
        pageList: [10, 25, 50, 100],        //可供选择的每页的行数（*）
        showRefresh: true,                  //是否显示刷新按钮
        search: true,                      //显示搜索框
        locale:"zh-CN",
        queryParams:{token:GetToken()},
        columns: [
            { field: 'id', title: 'ID',align:'center' },
            { field: 'itemname', title: '名称' ,align:'center' },
            { field: 'desc', title: '描述' ,align:'center' },
        ],
        responseHandler:function(res){
            if(res.code == 0){
                var temp =[]
                var data = res.data
                for(var i=0;i<data.length;i++){
                    temp.push({"id":data[i].id,"itemname":data[i].name,"desc":data[i].desc})
                }
                return temp
            }
            return []
        },
    });
}



function LoadPlayerTable(){
    $("#player-table").bootstrapTable("destroy");
    $("#player-table").bootstrapTable({
        url:"/player/allplayerinfo",
        method:"get",
        dataType:"json",
        pagination:true,
        clickToSelect: true,                //是否启用点击选中行
        sidePagination: "client",           //分页方式：client客户端分页，server服务端分页（*）
        pageNumber:1,                       //初始化加载第一页，默认第一页
        pageSize: 20,                       //每页的记录行数（*）
        pageList: [10, 25, 50, 100],        //可供选择的每页的行数（*）
        showRefresh: true,                  //是否显示刷新按钮
        search: true,                      //显示搜索框
        locale:"zh-CN",
        queryParams:{token:GetToken()},
        columns: [
            { field: 'id', title: 'ID' ,align:'center' },
            { field: 'name', title: '昵称' ,align:'center' },
            { field: 'silver', title: '银两' ,align:'center' },
            { field: 'diamond', title: '元宝' ,align:'center' },
            { field: 'lockstatus', title: '封号状态' ,align:'center' },
        ],
        responseHandler:function(res){
            if (res ==null || res.code!=0 || res.data ==null) return []
            if(res.code == 0){
                var temp =[]
                var data = res.data
                for(var i=0;i<data.length;i++){
                    temp.push({"id":data[i].id,"name":data[i].name,"silver":data[i].silver,"diamond":data[i].diamond,"lockstatus":data[i].lock_status})
                }
                return temp
            }
            return []
        },
    });
}


function LoadNoticeTable(){
    $("#notice-table").bootstrapTable("destroy");
    $("#notice-table").bootstrapTable({
        url:"/notice/allnotice",
        method:"get",
        dataType:"json",
        pagination:true,
        clickToSelect: true,                //是否启用点击选中行
        sidePagination: "client",           //分页方式：client客户端分页，server服务端分页（*）
        pageNumber:1,                       //初始化加载第一页，默认第一页
        pageSize: 20,                       //每页的记录行数（*）
        pageList: [10, 25, 50, 100],        //可供选择的每页的行数（*）
        showRefresh: true,                  //是否显示刷新按钮
        search: true,                      //显示搜索框
        locale:"zh-CN",
        queryParams:{token:GetToken()},
        columns: [
            { field: 'id', title: 'ID' ,align:'center' },
            { field: 'channel_id', title: '渠道号' ,align:'center' },
            { field: 'title', title: '公告标题' ,align:'center' },
            { field: 'content', title: '内容' ,align:'center' },
            { field: 'created', title: '创建时间' ,align:'center' },
            { field: 'starttm', title: '开始时间' ,align:'center' },
            { field: 'endtm', title: '结束时间' ,align:'center' },
        ],
        responseHandler:function(res){
            if (res ==null || res.code!=0 || res.data ==null) return []
            if(res.code == 0){
                var temp =[]
                var data = res.data
                for(var i=0;i<data.length;i++){
                    temp.push({"id":data[i].id,"channel_id":data[i].channel_id,"title":data[i].title,"content":data[i].content,"created":data[i].created_at,"starttm":data[i].start_time,"endtm":data[i].end_time})
                }
                return temp
            }
            return []
        },
    });
}




function LoadChannelTable(){
    $("#channel-table").bootstrapTable("destroy");
    $("#channel-table").bootstrapTable({
        url:"/channel/allchannel",
        method:"get",
        dataType:"json",
        pagination:true,
        clickToSelect: true,                //是否启用点击选中行
        sidePagination: "client",           //分页方式：client客户端分页，server服务端分页（*）
        pageNumber:1,                       //初始化加载第一页，默认第一页
        pageSize: 20,                       //每页的记录行数（*）
        pageList: [10, 25, 50, 100],        //可供选择的每页的行数（*）
        showRefresh: true,                  //是否显示刷新按钮
        search: true,                      //显示搜索框
        locale:"zh-CN",
        queryParams:{token:GetToken()},
        columns: [
            { field: 'id', title: 'ID' ,align:'center' },
            { field: 'name', title: '渠道名' ,align:'center' },
            { field: 'desc', title: '描述' ,align:'center' },
            { field: 'created', title: '创建时间' ,align:'center' },
        ],
        responseHandler:function(res){
            if (res ==null || res.code!=0 || res.data ==null) return []
            if(res.code == 0){
                var temp =[]
                var data = res.data
                for(var i=0;i<data.length;i++){
                    temp.push({"id":data[i].id,"name":data[i].name,"desc":data[i].desc,"created":data[i].created_at})
                }
                return temp
            }
            return []
        },
    });
}


function GetToken(){
    return Cookies.get("token")
}

function AddGoods(category,formname){
    var subdata = $("#"+formname).serialize() +"&token="+GetToken()+"&category="+category
    console.log(subdata)

    $.ajax({
        url:"/player/additem",
        async: true,
        type: "POST",
        data: subdata,
        beforeSend: function () {
            return true
        },
        success: function (res) {
            if (res.code !=0){
                alert("error:"+res.code)
            }
            //$('#add-item-result').text(res.code.toString())
            //alert("结果:"+ res.code.toString())
        },
    });
}

function AddRes(currency,formname){
    var subdata = $("#"+formname).serialize() +"&token="+GetToken()+"&currency="+currency
    console.log(subdata)
    $.ajax({
        url:"/player/addres",
        async: true,
        type: "POST",
        data: subdata,
        beforeSend: function () {
            return true
        },
        success: function (res) {
            if (res.code !=0){
                alert("error:"+res.code)
            }
        },
    });
}

function onClickedAddRewardItem(){
    var itemTypeID = $('#reward-type-id').val()
    var count = $('#reward-count').val()
    var category = $('input[name="optionsRadiosinline"]:checked').val()
    var str = '<li class="list-group-item" value="%s">%s<button onclick="$(this).parent().remove();return false">移除</button></li>'
    
    var item={Type:itemTypeID,Cate:category,Num:count}

    str = str.format(JSON.stringify(item),JSON.stringify(item))
    $("#reward-list-panel").append(str)
}

function SendMailSubmit(){

    var items = new Array();
    $("#reward-list-panel").children().each(function(i,n){
        var str = $(n).text().replace("移除","")
        var item = JSON.parse(str)
        items.push(item)
    });

    var liststr = JSON.stringify(items)
    var subdata = $("#send-mail-form").serialize() + "&rewardList="+liststr+ "&token="+GetToken()
    console.log(subdata)
    $.ajax({
        url:"/mail/sendmail",
        async: true,
        type: "POST",
        data: subdata,
        beforeSend: function () {
            return true
        },
        success: function (res) {
            if (res.code !=0){
                alert("error:"+res.code)
            }
        },
    });
}

function SendNoticeSubmit() {
    $.ajax({
        url:"/notice/sendnotice",
        async: true,
        type: "POST",
        data: $('#add-notice-form').serialize()+"&token="+GetToken(),
        beforeSend: function () {
            return true
        },
        success: function (res) {
            if (res.code !=0){
                alert("error:"+res.code)
            }
        },
    });
}

function RemoveNoticeSubmit(){
    $.ajax({
        url:"/notice/delnotice",
        async: true,
        type: "POST",
        data: $('#remove-notice-form').serialize()+"&token="+GetToken(),
        beforeSend: function () {

            return true
        },
        success: function (res) {
            if (res.code !=0){
                alert("error:"+res.code)
            }
        },
    });
}

function AddChannel(){
    $.ajax({
        url:"/channel/addchannel",
        async: true,
        type: "POST",
        data: $('#add-channel-form').serialize()+"&token="+GetToken(),
        beforeSend: function () {

            return true
        },
        success: function (res) {
            if (res.code !=0){
                alert("error:"+res.code)
            }
        },
    });
}

function DelChannel(){
    $.ajax({
        url:"/channel/delchannel",
        async: true,
        type: "POST",
        data: $('#remove-channel-form').serialize()+"&token="+GetToken(),
        beforeSend: function () {

            return true
        },
        success: function (res) {
            if (res.code !=0){
                alert("error:"+res.code)
            }
        },
    });
}



function hearthandler(){
    $.ajax({
        url:"/user/heart",
        async: true,
        type:"GET",
        data:{token:Cookies.get("token")},
        success:function(res){
            if (res.code!=0){
                var bret = confirm("token expired ! please reload!");
                toLogin();
            }
        },
        error:function (xhr) {
            var bret = confirm("heart error!");
            toLogin();
        }
    });
}

function resulthandler(){
    $.ajax({
        url:"/user/getresult",
        async: true,
        type:"GET",
        data:{token:Cookies.get("token")},
        success:function(res){
            if (res.code ==0 && res.data !=null){
                console.log(res)
                ClearResultInfo()
                for(var i=0;i<res.data.length;i++){
                    ShowResultInfo(res.data[i])
                }
            }
        }
    });
}

function LogoutClicked(){
    $.get("/user/logout",{"token":Cookies.get("token")},function (res) {
        toLogin();
    })
}

function toLogin(){
    Cookies.remove("token");
    location.href = "/login";
}


function ShowResultInfo(info){
    var temp = `
        <li>
            <div class="alert alert-%s">
                <a href="#" class="close" data-dismiss="alert">&times;</a>
                <strong>%s</strong>
            </div>
        </li>
    `
    if (info.indexOf("成功")!=-1){
        temp = temp.format("success",info)
    }else{
        temp = temp.format("warning",info)
    }
    $('#result-info-panel').append(temp)
}

function ClearResultInfo(){
    $('#result-info-panel').empty()
}


//------------------------------------------------------------------------------------------------------------------

String.prototype.format= function(){
    //将arguments转化为数组（ES5中并非严格的数组）
    var args = Array.prototype.slice.call(arguments);
    var count=0;
    //通过正则替换%s
    return this.replace(/%s/g,function(s,i){
        return args[count++];
    });
}
    