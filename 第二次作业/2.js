function huge(a){
    switch(a){
        case 1:
            console.log("这是case1");
            break;
        case 2:
            console.log("这是case2");
            break;
        default:
            console.log("这是默认情况");
    }
}

for(var i = 1; i < 4; i ++){
    huge(i)
}