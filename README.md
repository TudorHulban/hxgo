# hxgo

Typed, server-side hypermedia DSL that compiles to HTMX-compatible HTML

## 1. Introduction

### A. What hxgo is

hxgo is a typed, server‑side hypermedia DSL implemented as a Go library.
It provides a structured way to build user interfaces by expressing HTML, behavior, and interaction semantics through pure Go code.
The DSL compiles these typed constructs into HTMX‑compatible HTML, enabling interactive, server‑driven applications without templates or handwritten JavaScript.
By treating UI construction as a language rather than a collection of helpers, hxgo offers a predictable, composable, and Go‑native approach to building modern web interfaces.

### B. Why it exists

hxgo exists to provide a Go‑native way to build interactive web interfaces without relying on templates, JavaScript frameworks, or string‑based hypermedia attributes. Traditional HTMX usage requires developers to mix HTML, JavaScript, and backend logic, scattering UI behavior across multiple layers and making testing dependent on external tools such as Mocha, Jasmine, or browser‑based harnesses.

By contrast, hxgo expresses UI structure and behavior as pure Go values, enabling SPA‑level interactivity without client‑side frameworks, hydration steps, or npm‑based toolchains. This design allows developers to test UI generation, hypermedia behavior, and component logic using the standard Go testing framework, without JavaScript runtimes, DOM emulators, or external tooling. The broader performance and architectural advantages of this model are detailed in The Economics of Hypermedia, which explains why server‑rendered HTML is efficient, scalable, and well‑suited to modern network conditions.

### C. The one‑sentence pitch

hxgo provides a typed, server‑side hypermedia DSL that compiles pure Go code into HTMX‑compatible HTML, eliminating templates and the npm ecosystem while relying only on minimal vanilla JavaScript, offering Go‑native testing, a significantly narrower attack surface, and extremely low loading times for single‑page applications.

### D. The mental model (typed HTML + typed hypermedia controls)

hxgo is designed around a typed mental model in which user interfaces are expressed as pure Go values rather than templates, strings, or JavaScript. HTML elements are constructed through typed functions, attributes are represented as typed values, and hypermedia behavior is encoded through structured hx‑controls rather than free‑form strings. This approach treats UI construction as a declarative language embedded in Go, where each element, attribute, and interaction is validated at compile time. The result is a system in which UI structure, behavior, and server‑side logic form a single, coherent model that is easy to reason about, straightforward to test, and free from the fragility of template parsing or client‑side scripting.

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

### C. Server‑side rendering

All UI is rendered on the server, producing HTMX‑compatible HTML fragments that the client can swap into the DOM.

No hydration step or client‑side runtime.  
Simple deployment model with static assets optional.  
Predictable lifecycle driven by server responses.   
Full control over rendering and state transitions.

### D. Zero JavaScript

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

## 4. Quick Example

### A minimal “Good morning World!”

#### Page

#### Button

#### Backend Endpoint

#### The resulting behavior

## 5. Installation

```sh
go get github.com/TudorHulban/hxgo
```

## 6. Core Concepts

### Nodes

### Elements

### Attributes

### Behavior controls (hx-*)

### Widgets

### Server handlers

## 7. HTML DSL

### Element constructors

### Text nodes

### Conditional rendering

### Composition patterns

## 8. Hypermedia Controls (hx-attributes)

### hx-post

### hx-get

### hx-swap

### hx-send

### hx-require

### hx-upload

### UI toggles

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

### hxgo vs HTMX templates

### hxgo vs React/SPA

### hxgo vs templ

## 13. Roadmap

### More widgets

### Typed forms

### Client runtime helpers

### Documentation site

## 14. Contributing

### How to contribute

### Coding style

### Module boundaries

## 15. License

AGPL‑3.0
