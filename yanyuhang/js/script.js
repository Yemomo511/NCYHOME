let img_1=1
function like_onclick(){
    let input=document.getElementById("encourage1");
    let thank=document.getElementById("Thank1");
    if (img_1%2!=0){
        input.src="img/喜欢_click.png";
        thank.innerHTML="+1"
        img_1+=1
    }else{
        input.src="img/喜欢_like.png";
        thank.innerHTML=""
        img_1+=1
    }
}

let img_2=1
function like_onclick2(){
    let input=document.getElementById("encourage2");
    let thank=document.getElementById("Thank2");
    if (img_2%2!=0){
        input.src="img/喜欢_click.png";
        thank.innerHTML="+1"
        img_2+=1
    }else{
        input.src="img/喜欢_like.png";
        thank.innerHTML=""
        img_2+=1
    }
}
function returnTop(){
    var fanTop=document.getElementById("return_img");
    document.documentElement.scrollTop = document.body.scrollTop =0;//当点击元素时，让页面的滚动距离为0.写两个是为了兼容。
}