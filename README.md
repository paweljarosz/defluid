# DEFLUID

**Defluid** is a premade asset to create a 2D fluid simulation based on metaballs. It includes render stuff and a special emitter object that creates metaballs with a special material to simulate viscosity. To simulate physics it uses a Defold built-in Box2D physics, although it can be changed with whatever system you would like to.

Set the Defluid script properties as you wish or even copy the Defluid game object and create your own metaballs - for example with different collision objects, etc.

---

### Defluid emitter properties:

* **Num of metaballs** - number of created at once game objects
* **Scale of metaballs** - uniform scale of a metaball game object
* **Emitter width** - width of the box in which the metaballs are created
* **Emitter height** - height of the box in which the metaballs are created
* **Dry Each After** - will “dry” - animate each object’s scale down and at the end delete it after each amount of time (in seconds, set to 0 to turn drying off)
* **Create at init** - if true - creates num_of_metaballs objects at the defluid object creation
* **Debug physics** - if true - turns on Defold built-in physics debug visualisation

### Messages:

* **activate** – create metaballs based on object’s properties
* **change** – change object’s properties: num_of_metaballs, emitter_width, emitter_height, dry_each_after
* **force** – apply one force vector to each metaball
* **dry** – “dry” fluid - animate each metaball’s scale down to zero and at the end delete that object
* **remove** – delete all metaballs instantly

Additionaly you can set up the property of *defluid_plane*:
**Fluid Color** - allows to change the color of metaballs

and send message to *defluid_plane* to change the color of metaballs:
**change_color** – message.color [vmath.vector4()] - change a tint of rendered metaballs

---

### DEFLUID example
This example shows how can you modify defluid game object by sending messages to its script.

```
    function init(self)
        msg.post("/defluid", "activate")	-- activate creation of metaballs (fluid)
        msg.post("/defluid", "change", { dry_each_after = 1 })	-- change param of defluid
        timer.delay(3.5, false, function() 	-- after 3.5 s push up all metaballs
            msg.post("/defluid", "force", { force = vmath.vector3(0,10000000,0) })
        end)
        msg.post("/defluid", "dry")			-- activate drying (dry a metaball every second)
    end

    function final(self)
        msg.post("/defluid", "remove")		-- if needed you can destroy metaballs anytime
    end
```

#### To make it work:

* use included render with a custom render script (modify Render to /defluid/defluid.render in your game.project)
* add Rendercam library dependency in your game.project (I used https://github.com/rgrams/rendercam/archive/v1.0.1.zip 9)
* add Rendercam camera object to your collection
* add defluid_plane object (which has a render target plane) to your collection
* adjust your game.project to your needs - modify gravity, max collisions and max contacts (in example I used respectively -1000, 4096, 4096)
* add defluid object(s) to your collection and adjust its properties

*Note:* Defluid metaballs have its default collision objects with:
group: defluid
masks: defluid, ground, mob

Enjoy it and create awesome games with Defluid! :wink:

---

## Water Sources

I've made a simple game using this asset. It is available to play here: [https://paweljarosz.itch.io/water-sources](https://paweljarosz.itch.io/water-sources)
The source code is also available here: [https://github.com/paweljarosz/water-sources](https://github.com/paweljarosz/water-sources)

The game's aim is to flood the logo of Defold with a water. You can do this by tapping/clicking on the screen to explode a bomb. You have limited amount of bombs. And each droplet is drying in a moment. Have fun while playing with simulated fluid!

It also introduces some changes to Defluid (some physiscs adjustments and modified rendering: water surface is brighter than water depths)

---