### state paradigm

`stateManager` is an object passed down to all components, as an alternative to a redux store, passed via context. The `stateManager` has React's `state` and `setState` from `React.useState()` in closure.

`stateManager` contains:

- `state` for reading state.

- `ops` an operations object w/ methods. "Operations" are methods that induce application behavior, like clicking a button to get data & update state. Ideally, all event handlers are an operation inside this object. Operations are defined in `state/operations`.

- `changes` an object w/ methods for direct state changes. The idea is that these should be called by `operation` functions, and shouldn't have to be called directly. They are exposed here just in case escape is necessary, or if state needs to change in a very simple manner, and using an operation inside `ops` is too much overhead. State changes are defined in `state/stateChanges`.
