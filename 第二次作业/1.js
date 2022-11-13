var obj = {}
var obj1 = {
    uname: '张三疯',
        age:18,
            sex: '男',
                sayHi: function() {
                    console.log('hi~');
                }
}

function Star(uname, age, sex) {
    this.name = uname;
    this.age = age;
    this.sex = sex;
    this.sing = function(sang) {
        console.log(sang);

    }
}