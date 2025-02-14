<script lang="ts">
    import { onDestroy, onMount } from "svelte";

    let { roomData, mode, mouse = $bindable() }: { roomData: Record<string, {x: number, y: number}>, 
                                     mode: "client" | "server",
                                     mouse: {x: number, y: number} | null
                                   } = $props()

    class Ball {
        memberID: string
        x: number
        y: number
        radius: number
        color: string
        xvel: number
        yvel: number
        acc: number
        xtarget: number
        ytarget: number
        constructor(memberID: string, x: number, y: number, radius: number, color: string, 
                    xvel: number, yvel: number, acc: number, xtarget: number, ytarget: number) {
            this.memberID = memberID
            this.x = x
            this.y = y
            this.radius = radius
            this.color = color
            this.xvel = xvel
            this.yvel = yvel
            this.acc = acc
            this.xtarget = xtarget
            this.ytarget = ytarget
        }
        
        draw() {
            drawCircle(this.x, this.y, this.radius, this.color)
        }

        move() {
            this.xvel = (this.xtarget - this.x) / 5
            this.yvel = (this.ytarget - this.y) / 5

            this.x += this.xvel
            this.y += this.yvel
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
        for (const [memberID, pos] of Object.entries(roomData)) {
            const existingBall = ballList.find(ball => ball.memberID === memberID)
            if (existingBall) {
                existingBall.xtarget = pos.x
                existingBall.ytarget = pos.y
            } else {
                const newBall = new Ball(memberID, 50, 50, 50, getRandomColor(), 0, 0, 0, 50, 50)
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
        var yPos = (canvas.height / 100) * y
        ctx.fillStyle = color
        ctx.globalCompositeOperation = "color-burn"
        ctx.beginPath()
        ctx.ellipse(xPos,yPos,30,30,0,2*Math.PI,0)
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
            ball.move()
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

    function handleMouseMove(e: MouseEvent) {
        if (canvas && mouse) {
            var cRect = canvas.getBoundingClientRect(); // Gets CSS pos, and width/height
            
            const mouseX = e.clientX - cRect.left
            const mouseY = e.clientY - cRect.top

            mouse = {
                x: Math.round((mouseX / canvas.width) * 100), // Subtract the 'left' of the canvas
                y: Math.round((mouseY / canvas.height) * 100) // from the X/Y positions to make (0,0) the top left of the canvas
            }
        }            
    }

    // Track mouse movements if client
    onMount(() => {
        if (mode === "client" && canvas) {
            canvas.addEventListener("mousemove", handleMouseMove);
        }
    })

    onDestroy(() => {
        canvas?.removeEventListener("mousemove", handleMouseMove)
    })

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