# WebsocketGateway

WebSocket Gateway service built with Elixir Phoenix for the game engine platform.

## Features

- **Phoenix Channels** for real-time WebSocket communication
- **JWT Authentication** with token validation
- **Presence Tracking** for online users
- **Room Management** for game rooms
- **NATS Integration** for event broadcasting
- **Redis Integration** for caching and rate limiting

## Configuration

Environment variables:

- `WS_PORT` - WebSocket port (default: 8084)
- `WSS_PORT` - Secure WebSocket port (default: 8085)
- `JWT_SECRET_KEY` - JWT signing secret
- `REDIS_HOST` - Redis host
- `NATS_HOST` - NATS host
- `DB_HOST` - PostgreSQL host

## Running

```bash
# Development
mix deps.get
mix phx.server

# Production (with Docker)
docker build -t websocket-gateway .
docker run -p 8084:8084 -p 8085:8085 websocket-gateway
```

## Channels

- `game:{game_id}` - Game room channel
- `chat:{room_id}` - Chat channel
- `tournament:{tournament_id}` - Tournament channel
- `lobby` - Lobby channel

## License

MIT
