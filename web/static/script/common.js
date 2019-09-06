function Guid() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
        var r = Math.random()*16|0, v = c == 'x' ? r : (r&0x3|0x8);
        return v.toString(16);
    });
}

function FormatGoTime(t){
    var idx = t.indexOf(".")
    if (idx<0){
        idx = t.indexOf("+")
    }
    return t.substring(0,idx);
}

String.prototype.format = function() {
    var formatted = this;
    for (var i = 0; i < arguments.length; i++) {
        var regexp = new RegExp('\\{'+i+'\\}', 'gi');
        formatted = formatted.replace(regexp, arguments[i]);
    }
    return formatted;
};

function FormatStatus(s){
    if (s || s==="true"){
        return '<div><i class="layui-icon layui-icon-star-fill" style="color:#5FB878" ></i></div>';
    }
    return '';
}

function FormatConsulStatus(s){
    if (s==="passing"){
        return '<div><i class="layui-icon " style="color:#5FB878" >&#x1005;</i></div>';
    }
    return '<div><i class="layui-icon " style="color:#b8182a" >&#x1007;</i></div>';
}

function FormatSha(s){
    if (s.length <=12){
        return s;
    }
    if (s!=""){
        return '<span title="'+s+'">'+s.substring(0,4)+'...'+s.substring(s.length-8)+'</span>';
    }
    return "";
}

function FormatShaTip(sha,serviceSha) {
    if (sha == serviceSha){
        return '<span>'+FormatSha(sha)+'<i class="layui-icon" style="font-size: 15px; color: #FF5722;">&#xe756;</i></span>';
    }else{
        return FormatSha(sha)
    }
}

function FormatRunStatus(s) {
    switch (s){
        case 0: return "<i class='layui-icon layui-icon-ok-circle' style='color: #5fb878;'></i>";//"执行结束";
        case 1: return "<i class='layui-icon layui-icon-loading-1 layui-anim layui-anim-rotate layui-anim-loop' style='color: #3fa4ff;'></i>";//"执行中";
        case 2: return "<i class='layui-icon layui-icon-more' style='color: #91f5ff;'></i>";//"未执行";
        case 3: return "<i class='layui-icon layui-icon-ok-circle' style='color: #5fb878;'></i>";//"执行完";
        case 4: return "<i class='layui-icon layui-icon-about' style='color: #ff723d;'></i>";//"异常";
        case 5: return "<i class='layui-icon layui-icon-about' style='color: #fffd47;'></i>";//"后续异常";
    }
    return s;
}

function FormatRunKind(k,s) {
    if(s==2){
        return ""
    }

    if (k==0){
        return "无";
    }

    var result ="";
    var runKindBits= [
        ["2","更新程序"],
        ["4","更新配置"],
        ["8","更新其他"],
        ["4096","重载配置"],
        ["1","重启"],
        ["1024","上线"],
        ["2048","下线"]
    ];

    for(var i=0;i<runKindBits.length;i++){
        var bitArray = runKindBits[i];
        var bit = parseInt(bitArray[0]);
        if((bit&k) == bit){
            result+=bitArray[1]+" ";
        }
    }

    return result;
}

function OpenWinWidth() {
    var width =$(window).width();
    if (width<=1320){
        return (width-40)+'px';
    }else{
        return '1280px';
    }
}

function OpenWinHeight() {
    return ( $(window).height()-50)+'px';
}

String.prototype.format = function() {
    var formatted = this;
    for (var i = 0; i < arguments.length; i++) {
        var regexp = new RegExp('\\{'+i+'\\}', 'gi');
        formatted = formatted.replace(regexp, arguments[i]);
    }
    return formatted;
};

function ajaxSelect(url,data,id,selKey,form,emptyContent){
    $("#"+id).html("");
    $.post(url, data, function(result){
        if(result.code==0 && result.data){
            var opts = emptyContent;
            for(var i=0;i< result.data.length;i++){
                var key = result.data[i].Key;
                var value = result.data[i].Value;
                var selected ="";
                if(selKey == key){
                    selected = "selected";
                }
                opts += '<option value="{0}" {1}>{2}</option>'.format(key,selected,value);
            }
            $("#"+id).html(opts);
        }else{
            layer.msg(result.msg);
        }
        form.render();
    });
}

function hlTableTr(tableId,idx) {
    var trArr = $("[lay-id='"+tableId+"'] .layui-table-body.layui-table-main tr");
    if(trArr.length>0){
        hlRowTr($(trArr[idx]));
    }
    var trArr = $("[lay-id='"+tableId+"'] .layui-table-fixed.layui-table-fixed-l .layui-table-body tr");
    if(trArr.length>0){
        hlRowTr($(trArr[idx]));
    }
    var trArr = $("[lay-id='"+tableId+"'] .layui-table-fixed.layui-table-fixed-r .layui-table-body tr");
    if(trArr.length>0){
        hlRowTr($(trArr[idx]));
    }
}

function hlRowTr(rowTr) {
    rowTr.addClass('layui-btn-normal').siblings().removeClass('layui-btn-normal');
}