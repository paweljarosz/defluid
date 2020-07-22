embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/builtins/graphics/particle_blob.tilesource\"\n"
  "default_animation: \"anim\"\n"
  "material: \"/defluid/metaball.material\"\n"
  "blend_mode: BLEND_MODE_ADD\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 300.0\n"
  "friction: 0.0\n"
  "restitution: 0.6\n"
  "group: \"defluid\"\n"
  "mask: \"ground\"\n"
  "mask: \"defluid\"\n"
  "mask: \"mob\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 3.0\n"
  "}\n"
  "linear_damping: 0.1\n"
  "angular_damping: 1.0\n"
  "locked_rotation: true\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
