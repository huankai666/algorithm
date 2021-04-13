(function main(){
    var n = readline();
    var str = readline();
    var count = 0;
    var len = str.length;
    for(var i =0; i<len; i++){
        if(count==0&&str[i]=='L'){
            count = 3;
        }else if(count ==3&&str[i]=="R"){
            count = 0;
        }else if(str[i]=="L"){
            count --;
        }else if(str[i]=="R"){
            count ++;
        }
    }
    var res ;
    switch(count){
        case 0: res ="N"; break;
        case 1: res ="E"; break;
        case 2: res ="S"; break;
        case 3: res ="W"; break;
    }
    print(res);
})();