var contentAdminData;
var editContentAdminData;
/**
 * 管理员管理导航
 */
function adminNav() {
    navSelected('#adminNavItem');
    contentHeadShow('.adminContentHead');
    clearContentTable();
    initAdmin();
}

/**
 * 管理员导航初始化
 */
function initAdmin() {
    $('#contentData').bootstrapTable({
        data: contentAdminData,
        dataType: 'json',
        pagination: true,
        pageSize: 5,
        striped: true,
        search: false,
        singleSelect: false,
        showHeader: true,
        showFooter: false,
        showColumns: false,
        showRefresh: false,
        showToggle: false,
        sortable: false,
        columns: [{
            field: 'name',
            title: '账号',
            align: 'center',
            valign: 'middle'
        }, {
            field: 'password',
            title: '密码',
            align: 'center',
            valign: 'middle'
        }, {
            field: 'id',
            title: '操作',
            align: 'center',
            valign: 'middle',
            formatter: adminFormatter
        }]
    });
    initAdminData();
}

function initAdminData() {
    var token = localStorage.getItem("token")
    $.ajax({
        async: false,
        cache: false,
        type: 'GET',
        datType: "json",
        accept: "application/json;charset=UTF-8",
        contentType: "application/json;charset=UTF-8",
        url: 'http://localhost:3000/api/v1/admin',
        headers: {
            "Authorization": token
        },
        success: function (data) {
            if (data.code == 30005) {
                swal("请先登录！", data.msg, "error");
                window.location.href = "login.html";
                return
            }
            contentAdminData = data.data.adminList;
            var table = $('#contentData');
            table.bootstrapTable('refreshOptions', {data: contentAdminData, dataType: "json"});
        },
        error: function (data) {
        }
    });
}

function adminFormatter(value, row, index) {
    var id = value;
    var result = "";
    result += "<button type='button' class='btn btn-warning' data-toggle='modal' data-target='#adminUpdate' onclick=\"adminUpdate('" + index + "')\"><i class='fa fa-pencil'></i> 修改</button>";
    result += "<button type='button' class='btn btn-danger' onclick=\"adminDelete('" + id + "')\"><i class='fa fa-trash'></i> 删除</button>";
    return result;
}

/**
 * 管理员查询
 */
function adminQuery() {
    var token = localStorage.getItem("token")
    var adminName = $("#adminName").val();
    if (isNull(adminName)) {
        $.ajax({
            async: false,
            cache: false,
            type: 'GET',
            datType: "json",
            accept: "application/json;charset=UTF-8",
            contentType: "application/json;charset=UTF-8",
            url: 'http://localhost:3000/api/v1/admin',
            headers: {
                "Authorization": token
            },
            success: function (data) {
                if (data.code == 30005) {
                    swal("请先登录！", data.msg, "error");
                    window.location.href = "login.html";
                    return
                }
                contentAdminData = data._embedded.admin;
                var table = $('#contentData');
                table.bootstrapTable('refreshOptions', {data: contentAdminData, dataType: "json"});
            },
            error: function (data) {
            }
        });
    } else {
        $.ajax({
            async: false,
            cache: false,
            type: 'GET',
            datType: "json",
            accept: "application/json;charset=UTF-8",
            contentType: "application/json;charset=UTF-8",
            url: 'http://localhost:3000/api/v1/admin?name=' + adminName,
            headers: {
                "Authorization": token
            },
            success: function (data) {
                if (data.code == 30005) {
                    swal("请先登录！", data.msg, "error");
                    window.location.href = "login.html";
                    return
                }
                contentAdminData = dataArray;
                var table = $('#contentData');
                table.bootstrapTable('refreshOptions', {data: contentAdminData, dataType: "json"});
            },
            error: function (data) {
            }
        });
    }
}

function adminAddSave(){
    var token = localStorage.getItem("token")
    var data = {};
    data.name = $("#adminAddName").val();
    data.password = $("#adminAddPassword").val();
    $.ajax({
        async: true,
        cache: false,
        type: 'POST',
        data: JSON.stringify(data),
        datType: "json",
        accept: "application/json;charset=UTF-8",
        contentType: "application/json;charset=UTF-8",
        url: 'http://localhost:3000/api/v1/admin',
        headers: {
            "Authorization": token
        },
        success: function (data) {
            if (data.code == 30005) {
                swal("请先登录！", data.msg, "error");
                window.location.href = "login.html";
                return
            }
            swal('温馨提示', '新增管理员成功', 'success');
            initAdminData();
        },
        error: function (data) {
            swal('温馨提示', '新增管理员失败', 'error');
        }
    });
}

function adminUpdate(index) {
    var row = contentAdminData[index];
    editContentAdminData = contentAdminData[index];
    $("#adminUpdateId").val(row.id);
    $("#adminUpdateName").val(row.name);
    $("#adminUpdatePassword").val(row.password);
}

function adminUpdateSave(){
    var token = localStorage.getItem("token")
    var data = {};
    data.id = $("#adminUpdateId").val();
    data.name = $("#adminUpdateName").val();
    data.password = $("#adminUpdatePassword").val();
    data.createTime = editContentAdminData.createTime;
    $.ajax({
        async: true,
        cache: false,
        type: 'PUT',
        data: JSON.stringify(data),
        datType: "json",
        accept: "application/json;charset=UTF-8",
        contentType: "application/json;charset=UTF-8",
        url: 'http://localhost:3000/api/v1/admin',
        headers: {
            "Authorization": token
        },
        success: function (data) {
            if (data.code == 30005) {
                swal("请先登录！", data.msg, "error");
                window.location.href = "login.html";
                return
            }
            swal('温馨提示', '修改管理员成功', 'success');
            initAdminData();
        },
        error: function (data) {
            swal('温馨提示', '修改管理员失败', 'error');
        }
    });
}

function adminDelete(id) {
    var token = localStorage.getItem("token")
    $.ajax({
        async: true,
        cache: false,
        type: 'DELETE',
        datType: "json",
        accept: "application/json;charset=UTF-8",
        contentType: "application/json;charset=UTF-8",
        url: 'http://localhost:3000/api/v1/admin?id=' + id,
        headers: {
            "Authorization": token
        },
        success: function (data) {
            if (data.code == 30005) {
                swal("请先登录！", data.msg, "error");
                window.location.href = "login.html";
                return
            }
            swal('温馨提示', '删除管理员成功', 'success');
            initAdminData();
        },
        error: function (data) {
            swal('温馨提示', '删除管理员失败', 'error');
        }
    });
}

/**
 * 上传Excel文件
 */
function adminUpload() {
    var token = localStorage.getItem("token")
    var uploadData = new FormData();
    var uploadName = $("#adminUploadFile").val();
    uploadData.append("file", $("#adminUploadFile")[0].files[0]);
    uploadData.append("name", uploadName);
    $.ajax({
        url: 'http://localhost:3000/api/v1/upload?table=admin',
        headers: {
            "Authorization": token
        },
        type: 'POST',
        async: false,
        data: uploadData,
        // 告诉jQuery不要去处理发送的数据
        processData: false,
        // 告诉jQuery不要去设置Content-Type请求头
        contentType: false,
        beforeSend: function () {
            console.log("正在进行，请稍候");
        },
        success: function (data) {
            swal('温馨提示', '导入成功', 'success');
        }
    });
}