class QPathCreator {
    constructor(close) {
        this.points = []
        this.close = close
        this.fromPos = this.toPos = {x: 0, y: 0}
        this.started = false
        let ctrl = this
        qview.onmousedown = function(event) { ctrl.onmousedown(event) }
        qview.onmousemove = function(event) { ctrl.onmousemove(event) }
        qview.ondblclick = function(event) { ctrl.ondblclick(event) }
        qview.onkeydown = function(event) { ctrl.onkeydown(event) }
    }
    stop () {
        this.bezier()
        qview.onmousedown = null
        qview.onmousemove = null
        qview.ondblclick = null
        qview.onkeydown = null
    }

    async bezier() {
        var points = this.getPoints()
        var bezierStr =  await window.bezierPath(points);
        console.log(bezierStr);
        var bizierPath = document.getElementById("bezier-path")
        bizierPath.setAttribute("d", bezierStr);
        console.log(bizierPath);

        var svg_xml = (new XMLSerializer()).serializeToString(svg);
        var img = new Image();
        img.src = "data:image/svg+xml;base64," + window.btoa(svg_xml);
        img.onload = function () {
            //drawImage 可以用HTMLImageElement，HTMLCanvasElement或者HTMLVideoElement作为参数
            let ctx = document.getElementById("drawing").getContext('2d');
            ctx.drawImage(img, 0, 0);
        };
    }

    reset() {
        this.points = []
        this.started = false
        invalidate(null)
    }
    getPoints() {
        let points = [{x: this.fromPos.x, y: this.fromPos.y}]
        for (let i in this.points) {
            points.push(this.points[i])
        }
        return points;
    }
    buildShape() {
        let points = [{x: this.fromPos.x, y: this.fromPos.y}]
        for (let i in this.points) {
            points.push(this.points[i])
        }
        return new QPath(points, this.close, qview.lineStyle)
    }

    onmousedown(event) {
        this.toPos = qview.getMousePos(event)
        if (this.started) {
            this.points.push(this.toPos)
        } else {
            this.fromPos = this.toPos
            this.started = true
        }
        invalidate(null)
    }
    onmousemove(event) {
        if (this.started) {
            this.toPos = qview.getMousePos(event)
            invalidate(null)
        }
    }
    ondblclick(event) {
        if (this.started) {
            qview.doc.addShape(this.buildShape())
            this.reset()
        }
    }
    onkeydown(event) {
        switch (event.keyCode) {
        case 13: // keyEnter
            this.points.push(this.toPos)
            this.ondblclick(event)
            break
        case 27: // keyEsc
            this.reset()
        case 81:
            this.bezier()
        }
    }

    onpaint(ctx) {
        if (this.started) {
            let props = qview.properties
            ctx.lineWidth = props.lineWidth
            ctx.strokeStyle = props.lineColor
            ctx.beginPath()
            ctx.moveTo(this.fromPos.x, this.fromPos.y)
            for (let i in this.points) {
                ctx.lineTo(this.points[i].x, this.points[i].y)
            }
            ctx.lineTo(this.toPos.x, this.toPos.y)
            if (this.close) {
                ctx.closePath()
            }
            ctx.stroke()
        }
    }
}

qview.registerController("PathCreator", function() {
    return new QPathCreator(false)
})
