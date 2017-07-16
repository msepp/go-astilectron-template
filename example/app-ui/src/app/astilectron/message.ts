
// Request keys
export const REQUEST_WINDOW_CLOSE    = 'window.close';
export const REQUEST_WINDOW_MINIMIZE = 'window.minimize';
export const REQUEST_APP_VERSIONS    = 'get.versions';

// Event keys
export const EVENT_BACKEND_STATUS = 'backend.status';

// Message types
export const MESSAGE_TYPE_ALERT    = 'alert';
export const MESSAGE_TYPE_ERROR    = 'error';
export const MESSAGE_TYPE_EVENT    = 'event';
export const MESSAGE_TYPE_RESPONSE = 'response';

// Backend status keys
export const BACKEND_STATUS_IDLE = 'idle';

// BackendStatus describes backend status message
export class BackendStatus {
  constructor(
    public status?: string,
  ) {}
}

// Message describes an IPC message sent between render and backend process.
// Can be either a response, error (for a request) or an event.
export class Message {
  constructor(
    public id: string,
    public key: string,
    public type: string,
    public data: any
  ) {}
}
