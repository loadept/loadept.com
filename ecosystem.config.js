module.exports = {
  apps: [
    {
      name: "loadept.com",
      script: "app",
      interpreter: "none",
      instances: 1,
      autorestart: true,
      max_memory_restart: "200M",
      watch: false,
    }
  ]
}
