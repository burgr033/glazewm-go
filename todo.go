package main

/*

TODO: Fix the server stuff (messages should not be strings but Messages honoring message type)
TODO: check if the listening is even working
TODO: nothing in the subscription scope has been added yet
TODO: check what "binding_mode" and "focus_changed" should do in the JS
TODO: in the commands I need to honor the parameters (of course) like move window LEFT etc...

│   client.ts
└───types
    │   client-messages.ts
    │   server-message.ts
    │   wm-commands.ts
    │   wm-events.ts
    ├───events
    │       application-exiting-event.ts
    │       binding-mode-changed-event.ts
    │       focus-changed-event.ts
    │       focused-container-moved-event.ts
    │       monitor-added-event.ts
    │       monitor-removed-event.ts
    │       tiling-direction-changed-event.ts
    │       user-config-reloaded-event.ts
    │       window-managed-event.ts
    │       window-unmanaged-event.ts
    │       working-area-resized-event.ts
    │       workspace-activated-event.ts
    │       workspace-deactivated-event.ts

*/
