const BGS_HOST = `${window.location.protocol}//${window.location.hostname}:${
  window.location.hostname === "localhost" || "jaz1"
    ? "2470"
    : window.location.port
}`;

export { BGS_HOST };
