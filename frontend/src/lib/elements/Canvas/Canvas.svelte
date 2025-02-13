<script lang="ts">
    import { onMount } from "svelte";

    let { roomData }: { roomData: Record<string, number>} = $props()

    class Ball {
        memberID: string
        x: number
        y: number
        radius: number
        color: string
        xvel: number
        yvel: number
        acc: number
        constructor(memberID: string, x: number, y: number, radius: number, color: string, xvel: number, yvel: number, acc: number) {
            this.memberID = memberID
            this.x = x
            this.y = y
            this.radius = radius
            this.color = color
            this.xvel = xvel
            this.yvel = yvel
            this.acc = acc
        }
        
        draw() {
            drawCircle(this.x, this.y, this.radius, this.color)
        }
    }

    function getRandomColor() {
        var letters = '0123456789ABCDEF';
        var color = '#';
        for (var i = 0; i < 6; i++) {
            color += letters[Math.floor(Math.random() * 16)];
        }
        return color;
    }


    var ballList: Ball[] = $state([])

    $effect(() => {
        for (const [memberID, x] of Object.entries(roomData)) {
            const existingBall = ballList.find(ball => ball.memberID === memberID)
            if (existingBall) {
                existingBall.x = x
            } else {
                const newBall = new Ball(memberID, x, 50, 50, getRandomColor(), 0, 0, 0)
                ballList.push(newBall)
            }
        }
    })

    let canvas = $state<HTMLCanvasElement>()
    var ctx: CanvasRenderingContext2D | null = null

    function drawCircle(x: number, y: number, radius: number, color: string) {
        if (!ctx || !canvas) {
            return
        }
        var xPos = (canvas.width / 100) * x
        ctx.fillStyle = color
        ctx.globalCompositeOperation = "color-burn"
        ctx.beginPath()
        ctx.ellipse(xPos,30,30,30,0,2*Math.PI,0)
        ctx.fill()
        ctx.closePath()
    }


    $effect(() => {
        resizeCanvas()
    })

    function draw() {
        if (!ctx || !canvas) {
            return
        }
        ctx.clearRect(0, 0, canvas.width, canvas.height)

        ballList.forEach((ball) => {
            ball.draw()
        })

        requestAnimationFrame(draw)
    }

    onMount(() => {
        if (canvas) {
            ctx = canvas.getContext("2d")
        }

        resizeCanvas()

        requestAnimationFrame(draw)

        window.addEventListener('resize', () => {
            resizeCanvas()
        })
    })

    function resizeCanvas() {
        if (canvas) {
            canvas.width = document.body.clientWidth
            canvas.height = window.innerHeight
        }
    } 
</script>

<div id="canvas_container">
    <canvas id="canvas" bind:this={canvas}></canvas>
</div>



<style>
    #canvas_container {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin: 0px;
        padding: 0px;
    }

    #canvas {
        outline: 2px solid black;
        margin: 0px;
        padding: 0px;
    }
</style>