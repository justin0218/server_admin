package controllers

import (
	"github.com/astaxie/beego/context"
)

func Test (ctx *context.Context) {


	ctx.Output.Body([]byte(`<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <title>抽奖</title>
    <style>
        html,body{
            width: 100%;
            height: 100%;
        }
        body {
            background-color:#000000;
            color:#555555;
            overflow: hidden;
        }
        #canvasOne{
            display: block;
            margin: 0 auto;
            /*position: fixed;*/
            /*left: 0;*/
            /*right: 0;*/
            /*top: 0;*/
            /*bottom: 0;*/
            /*margin: auto;*/
            /*!*width: 100%;*!*/
            height: 100%;
        }
    </style>
</head>
<body>
<canvas id="canvasOne"></canvas>
<script type="text/javascript">
    function getLength(str){
        if(/[a-z]/i.test(str)){
            return str.match(/[a-z]/ig).length;
        }
        return 0;
    }

    window.addEventListener("load", init, false);
    var sphereRad = 480;
    var radius_sp = 1;
    var opt_display_dots = false;
    var unicodeFlakes = JSON.parse(localStorage.getItem("names")) || [];
    var timer;
    var theCanvas = document.getElementById("canvasOne");
    var context = theCanvas.getContext("2d");

    var displayWidth;
    var displayHeight;

    var wait;
    var count;
    var numToAddEachFrame;
    var particleList;
    var recycleBin;
    var particleAlpha;
    var r,g,b;
    var fLen;
    var m;
    var projCenterX;
    var projCenterY;
    var zMax;
    var turnAngle;
    var turnSpeed;
    var sphereCenterX, sphereCenterY, sphereCenterZ;
    var particleRad;
    var zeroAlphaDepth;
    var randAccelX, randAccelY, randAccelZ;
    var gravity;
    var rgbString;
    //we are defining a lot of variables used in the screen update functions globally so that they don't have to be redefined every frame.
    var p;
    var outsideTest;
    var nextParticle;
    var sinAngle;
    var cosAngle;
    var rotX, rotZ;
    var depthAlphaFactor;
    var i;
    var theta, phi;
    var x0, y0, z0;

    var grandPrize = localStorage.getItem("grand_prize");
    var firstPrize = localStorage.getItem("first_prize");
    var secondPrize = localStorage.getItem("second_prize");
    var thirdPrize = localStorage.getItem("third_prize");
    var fourPrize = localStorage.getItem("four_prize");


    var lucker = {}
    var lucker2 = [];

    function init() {
        wait = 1;
        count = wait - 1;
        numToAddEachFrame = 4;

        //particle color
        r = 70;
        g = 255;
        b = 140;

        rgbString = "rgba("+r+","+g+","+b+","; //partial string for color which will be completed by appending alpha value.
        particleAlpha = 1; //maximum alpha

        theCanvas.width = document.body.clientWidth;
        theCanvas.height = document.body.clientHeight;
        displayWidth = theCanvas.width;
        displayHeight = theCanvas.height;

        fLen = 520; //represents the distance from the viewer to z=0 depth.

        //projection center coordinates sets location of origin
        projCenterX = document.body.clientWidth/2 - 50;
        projCenterY = theCanvas.height/2;

        //we will not draw coordinates if they have too large of a z-coordinate (which means they are very close to the observer).
        zMax = fLen-2;

        particleList = {};
        recycleBin = {};

        //random acceleration factors - causes some random motion
        randAccelX = 0.1;
        randAccelY = 0.1;
        randAccelZ = 0.1;

        gravity = -0; //try changing to a positive number (not too large, for example 0.3), or negative for floating upwards.

        particleRad = 2.5;

        sphereCenterX = 0;
        sphereCenterY = 0;
        sphereCenterZ = -3 - sphereRad;

        //alpha values will lessen as particles move further back, causing depth-based darkening:
        zeroAlphaDepth = -750;

        turnSpeed = 2*Math.PI/2200; //the sphere will rotate at this speed (one complete rotation every 1600 frames).
        turnAngle = 0; //initial angle

    }

    document.onkeydown = function(e){
        // console.log(e.keyCode,e.which)
        if(!e){
            e = window.event;
        }
        if((e.keyCode || e.which) == 48){ //键盘0，特等抽取1位
            if(grandPrize){
                context.clearRect(0,0,displayWidth,displayHeight);
                context.font="60px Verdana";
                context.fillStyle = "rgb(70,255,140)";
                var engNameLen = getLength(grandPrize)/2;
                var m = grandPrize.length - getLength(grandPrize) + engNameLen;
                context.fillText("特等奖",displayWidth/2 - ((3 * 60)/2),displayHeight/2 - 120);
                context.fillText(grandPrize,displayWidth/2 - ((m * 60)/2),displayHeight/2 - 30);
                return
            }

            if(!timer){
                timer = setInterval(onTimer, 10/24);
                sphereRad = 680;
            }else{
                var t = setInterval(function(){
                    sphereRad -= 3;
                    if(sphereRad < 50){
                        clearInterval(t);
                        clearInterval(timer);
                        timer = null;
                        context.clearRect(0,0,displayWidth,displayHeight);

                        context.font="60px Verdana";
                        context.fillStyle = "rgb(70,255,140)";
                        var engNameLen = getLength(lucker.name)/2;
                        var m = lucker.name.length - getLength(lucker.name) + engNameLen;
                        context.fillText("特等奖",displayWidth/2 - ((3 * 60)/2),displayHeight/2 - 120);
                        context.fillText(lucker.name,displayWidth/2 - ((m * 60)/2),displayHeight/2 - 30);

                        for(var i = 0;i < unicodeFlakes.length;i++){ //删除
                            var item = unicodeFlakes[i];
                            if(item == lucker.name){
                                unicodeFlakes.splice(i,1);
                            }
                        }
                        localStorage.setItem("names",JSON.stringify(unicodeFlakes));
                        localStorage.setItem("grand_prize",lucker.name);
                        grandPrize = lucker.name;
                    }
                });
            }
        }

        if((e.keyCode || e.which) == 49){ ////键盘1，一等奖，抽取1位

            if(firstPrize){
                context.clearRect(0,0,displayWidth,displayHeight);
                context.font="60px Verdana";
                context.fillStyle = "rgb(70,255,140)";
                var engNameLen = getLength(firstPrize)/2;
                var m = firstPrize.length - getLength(firstPrize) + engNameLen;
                context.fillText("一等奖",displayWidth/2 - ((3 * 60)/2),displayHeight/2 - 120);
                context.fillText(firstPrize,displayWidth/2 - ((m * 60)/2),displayHeight/2 - 30);
                return
            }

            if(!timer){
                timer = setInterval(onTimer, 10/24);
                sphereRad = 680;
            }else{
                var t = setInterval(function(){
                    sphereRad -= 3;
                    if(sphereRad < 50){
                        clearInterval(t);
                        clearInterval(timer);
                        timer = null;
                        context.clearRect(0,0,displayWidth,displayHeight);

                        context.font="60px Verdana";
                        // 用渐变填色
                        context.fillStyle = "rgb(70,255,140)";
                        var engNameLen = getLength(lucker.name)/2;
                        var m = lucker.name.length - getLength(lucker.name) + engNameLen;
                        context.fillText("一等奖",displayWidth/2 - ((3 * 60)/2),displayHeight/2 - 120);
                        context.fillText(lucker.name,displayWidth/2 - ((m * 60)/2),displayHeight/2 - 30);
                        for(var i = 0;i < unicodeFlakes.length;i++){ //删除
                            var item = unicodeFlakes[i];
                            if(item == lucker.name){
                                unicodeFlakes.splice(i,1);
                            }
                        }
                        localStorage.setItem("names",JSON.stringify(unicodeFlakes));
                        localStorage.setItem("first_prize",lucker.name);
                        firstPrize = lucker.name;
                    }
                });
            }
        }

        if((e.keyCode || e.which) == 50){ ////键盘2，二等奖，抽取3位

            if(secondPrize){
                if(typeof secondPrize == "string"){
                    secondPrize = JSON.parse(secondPrize);
                }
                context.clearRect(0,0,displayWidth,displayHeight);
                context.font="60px Verdana";
                context.fillStyle = "rgb(70,255,140)";
                context.fillText("二等奖",displayWidth/2 - ((3 * 60)/2),displayHeight/2 - 180);

                var baseH = displayHeight/2 + 180;

                for(var i = 0;i < secondPrize.length;i++){
                    let luckers = secondPrize[i];
                    context.font="60px Verdana";
                    context.fillStyle = "rgb(70,255,140)";
                    var engNameLen = getLength(luckers)/2;
                    var m = luckers.length - getLength(luckers) + engNameLen;
                    baseH -= 90
                    context.fillText(luckers,displayWidth/2 - ((m * 60)/2),baseH);
                }
                return
            }

            if(!timer){
                timer = setInterval(onTimer, 10/24);
                sphereRad = 680;
            }else{
                var t = setInterval(function(){
                    sphereRad -= 3;
                    if(sphereRad < 50){
                        clearInterval(t);
                        clearInterval(timer);
                        timer = null;
                        context.clearRect(0,0,displayWidth,displayHeight);

                        context.font="60px Verdana";
                        context.fillStyle = "rgb(70,255,140)";
                        context.fillText("二等奖",displayWidth/2 - ((3 * 60)/2),displayHeight/2 - 180);

                        var baseH = displayHeight/2 + 180;


                        var secLs = [];
                        for(var i = 0;i < 3;i++){
                            let luckers = unicodeFlakes[Math.floor(Math.random()*unicodeFlakes.length)];
                            context.font="60px Verdana";
                            context.fillStyle = "rgb(70,255,140)";
                            var engNameLen = getLength(luckers)/2;
                            var m = luckers.length - getLength(luckers) + engNameLen;
                            baseH -= 90
                            context.fillText(luckers,displayWidth/2 - ((m * 60)/2),baseH);


                            for(var j = 0;j < unicodeFlakes.length;j++){ //删除
                                var item = unicodeFlakes[j];
                                if(item == luckers){
                                    unicodeFlakes.splice(j,1);
                                }
                            }
                            secLs.push(luckers);
                        }

                        localStorage.setItem("names",JSON.stringify(unicodeFlakes));
                        localStorage.setItem("second_prize",JSON.stringify(secLs));
                        secondPrize = JSON.stringify(secLs);
                    }
                });
            }
        }


        if((e.keyCode || e.which) == 51){ ////键盘3，三等奖，抽取5位

            if(thirdPrize){
                if(typeof thirdPrize == "string"){
                    thirdPrize = JSON.parse(thirdPrize);
                }
                context.clearRect(0,0,displayWidth,displayHeight);
                context.font="60px Verdana";
                context.fillStyle = "rgb(70,255,140)";
                context.fillText("三等奖",displayWidth/2 - ((3 * 60)/2),displayHeight/2 - 260);

                var baseH = displayHeight/2 + 280;

                for(var i = 0;i < thirdPrize.length;i++){
                    let luckers = thirdPrize[i];
                    context.font="60px Verdana";
                    context.fillStyle = "rgb(70,255,140)";
                    var engNameLen = getLength(luckers)/2;
                    var m = luckers.length - getLength(luckers) + engNameLen;
                    baseH -= 90
                    context.fillText(luckers,displayWidth/2 - ((m * 60)/2),baseH);
                }
                return
            }


            if(!timer){
                timer = setInterval(onTimer, 10/24);
                sphereRad = 680;
            }else{
                var t = setInterval(function(){
                    sphereRad -= 3;
                    if(sphereRad < 50){
                        clearInterval(t);
                        clearInterval(timer);
                        timer = null;
                        context.clearRect(0,0,displayWidth,displayHeight);

                        context.font="60px Verdana";
                        context.fillStyle = "rgb(70,255,140)";
                        context.fillText("三等奖",displayWidth/2 - ((3 * 60)/2),displayHeight/2 - 260);

                        var thirdPrizeArr = [];
                        var baseH = displayHeight/2 + 280;
                        for(var i = 0;i < 5;i++){
                            let luckers = unicodeFlakes[Math.floor(Math.random()*unicodeFlakes.length)];
                            context.font="60px Verdana";
                            context.fillStyle = "rgb(70,255,140)";
                            var engNameLen = getLength(luckers)/2;
                            var m = luckers.length - getLength(luckers) + engNameLen;
                            baseH -= 90
                            context.fillText(luckers,displayWidth/2 - ((m * 60)/2),baseH);


                            for(var j = 0;j < unicodeFlakes.length;j++){ //删除
                                var item = unicodeFlakes[j];
                                if(item == luckers){
                                    unicodeFlakes.splice(j,1);
                                }
                            }

                            thirdPrizeArr.push(luckers);
                        }


                        localStorage.setItem("names",JSON.stringify(unicodeFlakes));
                        localStorage.setItem("third_prize",JSON.stringify(thirdPrizeArr));
                        thirdPrize = JSON.stringify(thirdPrizeArr);


                    }
                });
            }
        }

        if((e.keyCode || e.which) == 52){ ////键盘4，四等奖，抽取10位

            if(fourPrize){
                if(typeof fourPrize == "string"){
                    fourPrize = JSON.parse(fourPrize);
                }

                context.clearRect(0,0,displayWidth,displayHeight);

                context.font="60px Verdana";
                context.fillStyle = "rgb(70,255,140)";
                context.fillText("四等奖",displayWidth/2 - ((3 * 60)/2),displayHeight/2 - 260);

                var baseH = displayHeight/2 + 280;
                var baseH2 = displayHeight/2 + 280;
                for(var i = 0;i < fourPrize.length;i++){
                    let luckers = fourPrize[i];
                    context.font="60px Verdana";
                    context.fillStyle = "rgb(70,255,140)";
                    var engNameLen = getLength(luckers)/2;
                    var m = luckers.length - getLength(luckers) + engNameLen;

                    if(i % 2 == 0){
                        baseH -= 90;
                        context.fillText(luckers,displayWidth/2 - ((m * 60)/2) + 300,baseH);
                    }else{
                        baseH2 -= 90;
                        context.fillText(luckers,displayWidth/2 - ((m * 60)/2) - 300,baseH2);
                    }


                }
                return;

            }

            if(!timer){
                timer = setInterval(onTimer, 10/24);
                sphereRad = 680;
            }else{
                var t = setInterval(function(){
                    sphereRad -= 3;
                    if(sphereRad < 50){
                        clearInterval(t);
                        clearInterval(timer);
                        timer = null;
                        context.clearRect(0,0,displayWidth,displayHeight);

                        context.font="60px Verdana";
                        context.fillStyle = "rgb(70,255,140)";
                        context.fillText("四等奖",displayWidth/2 - ((3 * 60)/2),displayHeight/2 - 260);


                        var baseH = displayHeight/2 + 280;
                        var baseH2 = displayHeight/2 + 280;
                        var fourPrizeArr = [];
                        for(var i = 0;i < 10;i++){
                            let luckers = unicodeFlakes[Math.floor(Math.random()*unicodeFlakes.length)];
                            context.font="60px Verdana";
                            context.fillStyle = "rgb(70,255,140)";
                            var engNameLen = getLength(luckers)/2;
                            var m = luckers.length - getLength(luckers) + engNameLen;

                            if(i % 2 == 0){
                                baseH -= 90;
                                context.fillText(luckers,displayWidth/2 - ((m * 60)/2) + 300,baseH);
                            }else{
                                baseH2 -= 90;
                                context.fillText(luckers,displayWidth/2 - ((m * 60)/2) - 300,baseH2);
                            }

                            for(var j = 0;j < unicodeFlakes.length;j++){ //删除
                                var item = unicodeFlakes[j];
                                if(item == luckers){
                                    unicodeFlakes.splice(j,1);
                                }
                            }
                            fourPrizeArr.push(luckers);

                        }
                        localStorage.setItem("names",JSON.stringify(unicodeFlakes));
                        localStorage.setItem("four_prize",JSON.stringify(fourPrizeArr));
                        fourPrize = JSON.stringify(fourPrizeArr);

                    }
                });
            }
        }

    }


    function addParticle(x0,y0,z0,vx0,vy0,vz0) {
        var newParticle;
        var color;
        //check recycle bin for available drop:
        if (recycleBin.first != null) {
            newParticle = recycleBin.first;
            //remove from bin
            if (newParticle.next != null) {
                recycleBin.first = newParticle.next;
                newParticle.next.prev = null;
            }
            else {
                recycleBin.first = null;
            }
        }
        //if the recycle bin is empty, create a new particle (a new ampty object):
        else {
            newParticle = {};
        }

        //add to beginning of particle list
        if (particleList.first == null) {
            particleList.first = newParticle;
            newParticle.prev = null;
            newParticle.next = null;
        }
        else {
            newParticle.next = particleList.first;
            particleList.first.prev = newParticle;
            particleList.first = newParticle;
            newParticle.prev = null;
        }

        //initialize
        newParticle.x = x0;
        newParticle.y = y0;
        newParticle.z = z0;
        newParticle.velX = vx0;
        newParticle.velY = vy0;
        newParticle.velZ = vz0;
        newParticle.age = 0;
        newParticle.dead = false;

        // newParticle.font="68px Arial";

        newParticle.flake = unicodeFlakes[Math.floor(Math.random() * unicodeFlakes.length)];
        if (Math.random() < 0.5) {
            newParticle.right = true;
        }
        else {
            newParticle.right = false;
        }
        return newParticle;
    }

    function recycle(p) {
        //remove from particleList
        if (particleList.first == p) {
            if (p.next != null) {
                p.next.prev = null;
                particleList.first = p.next;
            }
            else {
                particleList.first = null;
            }
        }
        else {
            if (p.next == null) {
                p.prev.next = null;
            }
            else {
                p.prev.next = p.next;
                p.next.prev = p.prev;
            }
        }
        //add to recycle bin
        if (recycleBin.first == null) {
            recycleBin.first = p;
            p.prev = null;
            p.next = null;
        }
        else {
            p.next = recycleBin.first;
            recycleBin.first.prev = p;
            recycleBin.first = p;
            p.prev = null;
        }
    }

    function onTimer() {
        //if enough time has elapsed, we will add new particles.
        count++;
        if (count >= wait) {

            count = 0;
            for (i = 0; i < numToAddEachFrame; i++) {
                theta = Math.random()*2*Math.PI;
                phi = Math.acos(Math.random()*2-1);
                x0 = sphereRad*Math.sin(phi)*Math.cos(theta);
                y0 = sphereRad*Math.sin(phi)*Math.sin(theta);
                z0 = sphereRad*Math.cos(phi);

                //We use the addParticle function to add a new particle. The parameters set the position and velocity components.
                //Note that the velocity parameters will cause the particle to initially fly outwards away from the sphere center (after
                //it becomes unstuck).
                var p = addParticle(x0, sphereCenterY + y0, sphereCenterZ + z0, 0.002*x0, 0.002*y0, 0.002*z0);

                //we set some "envelope" parameters which will control the evolving alpha of the particles.
                p.attack = 50;
                p.hold = 50;
                p.decay = 100;
                p.initValue = 0;
                p.holdValue = particleAlpha;
                p.lastValue = 0;

                //the particle will be stuck in one place until this time has elapsed:
                p.stuckTime = 90 + Math.random()*20;

                p.accelX = 0;
                p.accelY = gravity;
                p.accelZ = 0;
            }
        }

        //update viewing angle
        turnAngle = (turnAngle + turnSpeed) % (2*Math.PI);
        sinAngle = Math.sin(turnAngle);
        cosAngle = Math.cos(turnAngle);

        //background fill
        context.fillStyle = "#000000";
        context.fillRect(0,0,displayWidth,displayHeight);

        //update and draw particles
        p = particleList.first;
        while (p != null) {
            //before list is altered record next particle
            nextParticle = p.next;

            //update age
            p.age++;

            //if the particle is past its "stuck" time, it will begin to move.
            if (p.age > p.stuckTime) {
                p.velX += p.accelX + randAccelX*(Math.random()*2 - 1);
                p.velY += p.accelY + randAccelY*(Math.random()*2 - 1);
                p.velZ += p.accelZ + randAccelZ*(Math.random()*2 - 1);

                p.x += p.velX;
                p.y += p.velY;
                p.z += p.velZ;
            }

            /*
            We are doing two things here to calculate display coordinates.
            The whole display is being rotated around a vertical axis, so we first calculate rotated coordinates for
            x and z (but the y coordinate will not change).
            Then, we take the new coordinates (rotX, y, rotZ), and project these onto the 2D view plane.
            */
            rotX =  cosAngle*p.x + sinAngle*(p.z - sphereCenterZ);
            rotZ =  -sinAngle*p.x + cosAngle*(p.z - sphereCenterZ) + sphereCenterZ;
            m =radius_sp* fLen/(fLen - rotZ);
            p.projX = rotX*m + projCenterX;
            p.projY = p.y*m + projCenterY;

            //update alpha according to envelope parameters.
            if (p.age < p.attack+p.hold+p.decay) {
                if (p.age < p.attack) {
                    p.alpha = (p.holdValue - p.initValue)/p.attack*p.age + p.initValue;
                }
                else if (p.age < p.attack+p.hold) {
                    p.alpha = p.holdValue;
                }
                else if (p.age < p.attack+p.hold+p.decay) {
                    p.alpha = (p.lastValue - p.holdValue)/p.decay*(p.age-p.attack-p.hold) + p.holdValue;
                }
            }
            else {
                p.dead = true;
            }

            //see if the particle is still within the viewable range.
            if ((p.projX > displayWidth)||(p.projX<0)||(p.projY<0)||(p.projY>displayHeight)||(rotZ>zMax)) {
                outsideTest = true;
            }
            else {
                outsideTest = false;
            }

            if (outsideTest||p.dead) {
                recycle(p);
            }else {
                //depth-dependent darkening
                depthAlphaFactor = (1-rotZ/zeroAlphaDepth);
                depthAlphaFactor = (depthAlphaFactor > 1) ? 1 : ((depthAlphaFactor<0) ? 0 : depthAlphaFactor);
                context.fillStyle = rgbString + depthAlphaFactor*p.alpha + ")";
                /*ADD TEXT function!*/
                /*ADD TEXT function!*/
                /*ADD TEXT function!*/
                /*ADD TEXT function!*/
                context.fillText(p.flake,p.projX, p.projY);

                lucker = {"name":p.flake};

                context.font="28px Arial";
                /*ADD TEXT function!*/
                /*ADD TEXT function!*/
                /*ADD TEXT function!*/
                /*ADD TEXT function!*/
                //draw
                context.beginPath();
                if(opt_display_dots){
                    context.arc(p.projX, p.projY, m*particleRad, 0, 2*Math.PI, false);
                }
                context.closePath();
                context.fill();
            }

            p = nextParticle;
        }
    }

</script>

</body>
</html>
`))
}
