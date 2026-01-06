# hxgo

Typed, server-side hypermedia domain specific language (DSL) that renders to HTMX-compatible HTML.  
hxgo treats HTML as a render target rather than a source language, using a typed Node graph as an intermediate representation that deterministically lowers Go values into hypermedia.  
hxgo incorporates concepts from [htmx](https://htmx.org) and [gomponents.com](https://www.gomponents.com/).

## 1. Introduction

### A. What hxgo is

hxgo is a typed, server‑side hypermedia typed UI language implemented as a Go library.
It provides a structured way to build user interfaces by expressing HTML, behavior, and interaction semantics through pure Go code.
The DSL renders these typed constructs into HTMX‑compatible HTML, enabling interactive, server‑driven applications without templates or handwritten JavaScript.
By treating UI construction as a language rather than a collection of helpers, hxgo offers a predictable, composable, and Go‑native approach to building modern web interfaces.

hxgo is not a template engine, a client-side framework, or a code generator. It does not interpret markup, parse templates, or generate Go source files. Instead, hxgo treats user interfaces as typed Go values and renders them directly into HTML at runtime with deterministic semantics and minimal allocation.

hxgo intentionally trades generality for correctness, predictability, and performance. It assumes that UI interactions originate from hxgo-generated HTML and that hypermedia contracts are respected. As a result, hxgo is not designed to be a general-purpose HTTP parsing framework or a drop-in replacement for defensive, open-world middleware.

These constraints allow hxgo to:

- minimize runtime checks
- reduce allocations
- simplify request parsing
- enforce stronger invariants through types

A request is handled by a server handler, which constructs a tree of typed Nodes. These Nodes are rendered in a single pass into HTML. HTMX interprets the hypermedia controls in that HTML and issues the next request. The server handler responds by constructing a new Node tree, rendering it to HTML, and returning it for client-side DOM replacement.

Request → Handler → Node(s) → Render → HTML → HTMX → Request → Handler → Node(s) → Render → HTML → HTMX swap

### B. Why it exists

hxgo exists to provide a Go‑native way to build interactive web interfaces without relying on templates, JavaScript frameworks, or string‑based hypermedia attributes. Traditional HTMX usage requires developers to mix HTML, JavaScript, and backend logic, scattering UI behavior across multiple layers and making testing dependent on external tools such as Mocha, Jasmine, or browser‑based harnesses.

By contrast, hxgo expresses UI structure and behavior as pure Go values, enabling SPA‑level interactivity without client‑side frameworks, hydration steps, or npm‑based toolchains. This design allows developers to test UI generation, hypermedia behavior, and component logic using the standard Go testing framework, without JavaScript runtimes, DOM emulators, or external tooling. The broader performance and architectural advantages of this model are detailed in The Economics of Hypermedia, which explains why server‑rendered HTML is efficient, scalable, and well‑suited to modern network conditions.

### C. The one‑sentence pitch

hxgo provides a typed, server‑side hypermedia embedded language that renders pure Go code into HTMX‑compatible HTML, eliminating templates and the npm ecosystem while relying only on minimal vanilla JavaScript, offering Go‑native testing, a significantly narrower attack surface, and extremely low loading times for single‑page applications.

### D. The mental model (typed HTML + typed hypermedia controls)

hxgo is designed around a typed mental model in which user interfaces are expressed as pure Go values rather than templates, strings, or JavaScript. HTML elements are constructed through typed functions, attributes are represented as typed values, and hypermedia behavior is encoded through structured hx‑controls rather than free‑form strings. This approach treats UI construction as a declarative language embedded in Go, where each element, attribute, and interaction is validated through typed construction, with illegal states unrepresentable. The result is a system in which UI structure, behavior, and server‑side logic form a single, coherent model that is easy to reason about, straightforward to test, and free from the fragility of template parsing or client‑side scripting.

hxgo follows a contract-driven hypermedia model: the server defines both the structure of the UI and the shape of the interactions that return to it. Requests are treated as continuations of server-emitted contracts rather than arbitrary inputs.

Rendering model

hxgo renders UI using a single-pass, depth-first traversal of a typed node tree. Each node contributes its HTML representation directly into a shared accumulator, producing a contiguous HTML byte stream with no intermediate buffers or interpretation steps.

Rendering has the following properties:

- Single-pass traversal
- No template execution or AST interpretation
- No virtual DOM
- Deterministic output for a given node tree
- Minimal allocations through pre-sized accumulators

Because rendering is a pure transformation from Nodes to HTML, the same rendering logic applies uniformly to full pages, fragments, and component updates.

### E. Design Constraints

hxgo is intentionally designed around a small set of explicit constraints. These constraints are not limitations in the traditional sense; they are architectural choices that enable determinism, performance, and strong guarantees about correctness.

#### 1. Closed-world UI construction

hxgo assumes a closed world for UI construction: all UI elements, attributes, hypermedia controls, and CSS rules are defined within the Go program at compile time.

There is no support for:

- runtime injection of arbitrary HTML,
- user-provided markup,
- dynamic evaluation of template strings.
- This constraint allows hxgo to:
- guarantee structural correctness of generated HTML,
- prevent malformed or unsafe hypermedia attributes,
- eliminate entire classes of injection vulnerabilities,
- reason about the complete UI graph at render time.

Because the UI surface is known ahead of time, hxgo can safely apply optimizations such as deduplication, deterministic ordering, and single-pass rendering.

#### 2. HTML as a render target, not a source language

hxgo does not parse, interpret, or manipulate HTML templates. Instead, HTML is treated strictly as a serialization format emitted by the renderer.

UI structure is expressed exclusively as typed Go values.
HTML is the output, never the input.

This design:

- avoids template parsing and runtime evaluation,
- removes ambiguity between data and markup,
- guarantees deterministic output for a given Node graph,
- enables compile-time validation of UI structure and behavior.

#### 3. Typed hypermedia semantics

All hypermedia interactions (hx-get, hx-post, swaps, requirements, enables/disables, uploads, redirects) are expressed as typed constructs.

Free-form string-based hypermedia attributes are deliberately avoided.

This constraint ensures:

- invalid or incomplete hypermedia configurations cannot be constructed,
- interaction semantics are consistent across the entire UI,
- hypermedia behavior is testable using standard Go tooling,
- server-side handlers and client-side behavior remain aligned.

#### 4. Deterministic rendering and ordering

Rendering in hxgo is deterministic by design. Given the same Node graph, the renderer will always emit identical HTML and CSS.

This is enforced through:

- fixed traversal order of Node graphs,
- stable attribute emission order,
- controlled CSS accumulation and deduplication,
- absence of reflection-based rendering.

Determinism enables:

- reliable snapshot testing,
- predictable caching behavior,
- reproducible builds and responses.

#### 5. Server-driven state transitions

hxgo does not model client-side state machines. All state transitions are driven by server responses and expressed through hypermedia.

Client-side JavaScript is limited to HTMX’s minimal runtime or optional, or declarative browser APIs (e.g. window.open).

This constraint:

- avoids hydration and reconciliation layers,
- eliminates client-side state divergence,
- keeps all authoritative state on the server,
- aligns with RESTful hypermedia principles.

#### 6. Controlled extensibility

Extensibility in hxgo is achieved through:

- new Node constructors,
- new typed attributes,
- new widgets composed from existing primitives.

The core renderer and accumulator are intentionally unexported. This prevents external code from mutating internal invariants and ensures that all extensions preserve correctness and determinism.

Advanced features such as CSS registries, media-query handling, and deduplication are implemented internally and exposed only through stable, typed APIs.

#### 7. Go-first developer experience

hxgo is designed for Go developers who prefer:

- static typing,
- explicit control flow,
- standard tooling (go test, go vet, pprof),
- minimal build pipelines.

As a result:

- no JavaScript tooling is required,
- no template files are introduced,
- UI logic can be tested without browsers or DOM emulators,
- performance characteristics are transparent and measurable.

## 2. The Economics of Hypermedia

Modern network conditions and server capabilities make hypermedia‑driven applications significantly more efficient than traditional JSON‑based REST or GraphQL architectures. Instead of serializing data into JSON, transmitting it, parsing it on the client, and reconstructing UI state, hxgo sends directly rendered HTML fragments, reducing both computational overhead and architectural complexity.

### A. Network efficiency

HTML fragments are typically small and compressible, often fitting within a 15 KB TCP/IP frame, making their transmission cost negligible compared to the latency and CPU overhead of JSON parsing and client‑side rendering.

Compression‑friendly HTML fragments reduce payload size dramatically.  
Single‑frame responses minimize round‑trip latency.  
Predictable payload structure avoids schema negotiation.

### B. Eliminating JSON serialization overhead

JSON‑based systems require reflection, encoding, decoding, schema validation, and client‑side parsing. Hypermedia avoids these steps entirely by rendering UI directly from server‑side objects.

No reflection‑based encoding.  
No client‑side parsing.  
No schema mismatch risks.  
No hydration or virtual DOM.

### C. Avoiding REST and GraphQL layers

Hypermedia removes the need for REST endpoints, DTOs, GraphQL resolvers, and client‑side state machines. The server returns the UI itself, not a data structure that must be interpreted.

No REST boilerplate, not too much or to little data.  
No GraphQL schema machinery.  
No client‑side state reconstruction.  
Direct object‑to‑HTML rendering.

### D. Faster application startup

Because hxgo applications avoid hydration, bundling, and client‑side frameworks, they start almost instantly, benefiting from Go’s native binary performance.

No npm toolchain.  
No JS bundling or transpilation.  
No hydration step.  
Minimal client runtime.

### E. Single‑frame CSS delivery via data URLs  

Embedding component‑scoped CSS directly into the HTML response using data:text/css;base64 URLs eliminates an entire network round‑trip and removes the need for external stylesheet management. Because both the HTML fragment and its associated CSS typically fit comfortably within a ~15 KB TCP/IP frame, the browser receives everything it needs to render the component in a single packet, without waiting for additional resources.

No external stylesheet fetch: CSS arrives inline with the HTML, eliminating render‑blocking requests.
Single‑frame delivery: HTML + base64‑encoded CSS often fit within one TCP frame, minimizing latency.
Deterministic styling: Each component carries its own isolated CSS, avoiding global leakage and cascade conflicts.
No asset pipeline: No bundling, no minification, no cache invalidation, no file management.
Immediate paint: The browser can render the component as soon as the first response arrives.

This approach restores hypermedia’s original strength: a complete, self‑contained document delivered atomically, without auxiliary resources or client‑side orchestration.  

The usual file can also be referred and the element will receive the CSS based on the classes set.  
Note that this method has a small performance penalty.  

## 3. Key Features

### A. Typed HTML DSL

hxgo provides a typed HTML construction layer, where every element, attribute, and node is represented as a Go value rather than a string. This eliminates template fragility and ensures that UI structure is validated at compile time.

Compile‑time validation of HTML structure.  
Pure Go expressions instead of templates.  
No string concatenation or manual markup assembly.  
Predictable rendering semantics based on typed nodes.

### B. Typed hx‑attributes

Hypermedia behavior is expressed through typed hx‑controls, ensuring that interactions such as hx-post, hx-get, and hx-swap are constructed safely and consistently.

Typed hypermedia controls instead of free‑form strings.  
Consistent behavior encoding across components.  
Compile‑time safety for all hx‑attributes.  
Elimination of malformed attributes and runtime surprises.

In string-based HTMX usage, invalid or contradictory attributes such as hx-get and hx-post on the same element are syntactically legal but semantically incorrect, and errors surface only at runtime. In hxgo, such invalid states cannot be constructed, because hypermedia interactions are encoded as typed values with enforced invariants.

### C. Server‑side rendering

All UI is rendered on the server, producing HTMX‑compatible HTML fragments that the client can swap into the DOM.

No hydration step or client‑side runtime.  
Simple deployment model with static assets optional.  
Predictable lifecycle driven by server responses.   
Full control over rendering and state transitions.

### D. Zero Custom JavaScript

hxgo removes the need for handwritten JavaScript, npm tooling, or browser‑side frameworks.

No npm ecosystem or dependency chains.  
No build pipelines for JS bundling.  
No client‑side state management.  
Reduced cognitive load for Go developers.

### E. Hypermedia interactions

UI behavior is encoded directly in the HTML via typed hx‑controls, enabling a hypermedia‑driven interaction model.

Server‑driven UI updates.  
Declarative behavior encoded in markup.  
HTMX‑compatible semantics.  
Minimal client footprint.  

### F. Composable widgets

hxgo supports reusable, composable UI components built entirely in Go.

Reusable components with typed inputs.  
Encapsulated behavior and rendering.  
Predictable composition through pure functions.  
No template partials or string injection.

### G. Secure by design

Typed HTML and typed hypermedia controls significantly reduce the attack surface.

No template injection.  
No inline JavaScript.  
No user‑controlled markup.  
Narrower attack surface through typed constructs.

### H. Typed request continuation

hxgo treats incoming requests as typed continuations of previously emitted hypermedia actions rather than as arbitrary payloads. Because forms and actions are generated by hxgo itself, request parsing can rely on narrower, well-defined contracts, enabling simpler parsing logic, fewer allocations, and predictable performance without sacrificing correctness.

### K. Memory Safety Guarantees

hxgo uses `unsafe.Pointer` internally but maintains memory safety through:

1. **Request-Scoped Construction**: Node trees are built and rendered within a single request handler
2. **No Cross-Goroutine Sharing**: Each Node tree is owned by exactly one goroutine
3. **Immediate Rendering**: Nodes are rendered immediately after construction
4. **Heap Allocation**: Escape analysis ensures all Node data lives on the heap

This combination allows performance optimizations while preserving safety.  
For caching do not store Nodes but HTML.

## 4. Quick Example

Nodes should not be shared across goroutines.  
Nodes can be created in a presentation layer, using below the handler for keeping code concise:

```go
// Request handler (single goroutine)
func handler(w http.ResponseWriter, r *http.Request) {
    // 1. Build fresh Node tree
    ui := el("div",
        el("p", Text("Hi")),
        el("p", Text("World!")),
    )
    
    // 2. Render immediately (same goroutine)
    html := Render(ui)
    
    // 3. Send response
    w.Write(html)
    // 4. ui goes out of scope → GC can collect
}
```

## 5. Installation

```sh
go get github.com/TudorHulban/hxgo
```

## 6. Core Concepts

hxgo is built on a small set of composable primitives. These primitives form the foundation of the language and define how UI structure, behavior, and composition are expressed in typed Go.

### Nodes

A Node is the fundamental unit of hxgo. Every piece of UI — elements, text, attributes, CSS, widgets — ultimately resolves to a Node. Nodes form a tree that mirrors the structure of the resulting HTML. It represents a typed intermediate representation executed via a deterministic accumulator, as opposed to an interface-based render method. This design enables CSS deduplication, predictable attribute order, and first-class hypermedia controls, which are beyond the scope of simple HTML rendering.

Unified representation: all UI constructs reduce to a Node.  
Tree‑structured output: Nodes compose into hierarchical HTML.  
Pure values: Nodes contain no side effects or rendering logic.  
Render‑ready: Nodes serialize deterministically into HTML or CSS.  
Nodes are the “atoms” of the language.

A Node becomes HTML through a simple, predictable process:

1. The Node exposes a render method that serializes its structure into a stable HTML representation.  
2. Composite Nodes recursively render their children, producing a hierarchical output.  
3. Attributes are rendered in a fixed order to ensure deterministic output.  

During this process the components CSS is accumulated and consolidated.

The final result is a byte slice or string containing valid HTML, ready to be sent to the client, together with the CSS the components need.  
This rendering model ensures that UI generation is pure, testable, and free from side effects.

#### CSS

CSS is accumulated as styles or CSS.  
Styles could be embedded directly in the request answer.  
CSS end goal is generating real `.css` files.  

### Elements

An Element is a Node that represents a HTML tag. Elements contain:

- a tag name
- zero or more attributes
- zero or more child nodes

Elements are constructed through typed functions such as Div(), Span(), Button(), and so on.  
Typed constructors: each HTML tag maps to a Go function.  
Structured children: Elements contain other Nodes.  
Deterministic rendering: output HTML is predictable.  
No string concatenation: structure is encoded, not assembled.

Elements define the structural skeleton of the UI.

### Attributes

An Attribute is a typed value representing a HTML attribute. Attributes attach metadata or behavior to Elements. Unlike string‑based HTML, hxgo ensures attributes are well‑formed and type‑checked.

Typed attribute values: no free‑form strings.  
Compile‑time safety: invalid attributes cannot be constructed.  
Consistent semantics: attributes behave uniformly across components.  
Composable sets: attributes can be grouped and reused.  
Attributes define the configuration of an Element.

### Behavior controls (hx-*)

Behavior controls represent hypermedia interactions such as `hx-get` or `hx-post`. In hxgo, these are expressed as typed constructs rather than raw strings.

Typed hypermedia controls: no malformed hx‑attributes.  
Declarative behavior: interactions encoded directly in markup.  
Server‑driven updates: behavior defined on the server.  
Predictable semantics: consistent across all components.  
Behavior controls define how the UI interacts with the server.

### Widgets

A Widget is a reusable, composable function that returns a Node. Widgets encapsulate structure, attributes, and behavior into a single unit that can be reused across the application.

Pure functions: widgets take data and return Nodes.  
Encapsulated behavior: structure and hx‑controls bundled together.  
Reusable components: composable across the UI.  
No templates: all logic expressed in typed Go.  

Widgets are the building blocks of higher‑level UI composition.

### Server handlers

A server handler is the bridge between hxgo’s typed UI and the application’s backend logic. Handlers receive HTTP requests, perform domain‑specific work by callng domain services, and return Nodes that hxgo renders into HTML fragments. These fragments are then swapped into the client’s DOM by HTMX’s minimal vanilla JavaScript runtime.

Handlers define how the server responds to hypermedia interactions, and they form the execution boundary between UI construction and application logic.

A handler performs the following responsibilities:

1. Accepts an HTTP request: receives form data, query parameters, or event triggers.
2. Executes validations and delegates domain logic: calls domain services to perform updates, queries, or computations when the interaction requires domain behavior.
3. Coordinates with the presentation layer:

    - for purely presentational pages (for example, a landing or login page), the handler can invoke the presentation layer directly to obtain a Node,
    - for domain‑driven interactions, the handler passes the resulting domain objects to the presentation layer, where they are injected into widget functions and transformed into a Node.

4. Renders the Node into HTML and writes the response: the Node is rendered into HTML and sent back to the client.

This model keeps the UI and backend unified in a single language and eliminates the need for JSON APIs, DTOs, or client‑side state machines.

## 7. HTML DSL

### Element constructors

### Text nodes

### Conditional rendering

### Composition patterns

## 8. Hypermedia Controls (hx-attributes)

See [HX Attributes](./js/README.md).

## 9. Widgets

### ButtonSubmit

### Forms

### Inputs

### Reusable components

## 10. Server Integration

### Endpoints

### Returning HTML fragments

### Swap strategies

### Redirects

## 11. Security Model

### Typed HTML

### No JS injection

### No template injection

### Narrow attack surface

## 12. Comparison

The comparisons below focus on architectural intent rather than surface syntax. hxgo occupies a different design space than template engines and client-side frameworks by modeling hypermedia interactions as typed server-side constructs rather than as strings or client-side state machines.

templ → “typed markup, string-based hypermedia”

React → “client-side state machine”

HTMX templates → “string-assembled hypermedia”

### hxgo vs HTMX templates

### hxgo vs React/SPA

### hxgo vs templ

## 13. Roadmap

### More widgets

### Typed forms

### Client runtime helpers

### Documentation site

## 14. Contributing

See [Contributing](Contributing.md).

## 15. License

AGPL‑3.0
