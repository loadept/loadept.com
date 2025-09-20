module.exports = {
  apps: [
    {
      name: "loadept-dev",
      script: "cmd/loadept/main.go",
      interpreter: "go",
      interpreter_args: "run",
      watch: ["cmd/", "api/", "internal/", "pkg/"],
      cwd: "./",
      restart_delay: 5000,
    }
  ]
}
