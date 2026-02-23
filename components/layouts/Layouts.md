# Layout System Documentation

## Core Philosophy

The folder provides a library of opinionated, named layouts. Each layout serves a specific content type and knows how to rearrange itself at different screen sizes.

## Layout 1: Documentation

Purpose: Content with navigation/supporting info (docs, tutorials, reference).

### 1.1. Documentation - Desktop (≥1024px)

```css
┌─────────────────────────────────┐
│            header?              │
├───────────────┬─────────────────┤
│               │                 │
│   sidebar     │     main        │
│   (250px)     │                 │
│               │                 │
├───────────────┴─────────────────┤
│            footer?              │
└─────────────────────────────────┘
```

### 1.2 Documentation - Mobile (<640px)

```css
┌─────────────────────────────────┐
│            header?              │
├─────────────────────────────────┤
│     menu button (if collapsed)  │
├─────────────────────────────────┤
│                                 │
│             main                │
│                                 │
├─────────────────────────────────┤
│            footer?              │
└─────────────────────────────────┘
```

## Layout 2: Dashboard

Purpose: Data overview, analytics, management interfaces.

### 2.1 Dashboard - Desktop (≥1024px)

```css
┌─────────────────────────────────┐
│             header              │
├─────────────────────────────────┤
│           filters?              │
├─────────────────────────────────┤
│ ┌─────┐ ┌─────┐ ┌─────┐         │
│ │card │ │card │ │card │         │
│ └─────┘ └─────┘ └─────┘         │
│ ┌─────┐ ┌─────┐ ┌─────┐         │
│ │card │ │card │ │card │         │
│ └─────┘ └─────┘ └─────┘         │
├─────────────────────────────────┤
│             footer              │
└─────────────────────────────────┘
```

### 2.2. Dashboard - Tablet (640px - 1023px)

```css
┌─────────────────────────────────┐
│             header              │
├─────────────────────────────────┤
│ ┌─────┐ ┌─────┐                 │
│ │card │ │card │                 │
│ └─────┘ └─────┘                 │
│ ┌─────┐ ┌─────┐                 │
│ │card │ │card │                 │
│ └─────┘ └─────┘                 │
└─────────────────────────────────┘
```

### 2.3 Dashboard - Mobile (<640px)

```css
┌─────────────────────────────────┐
│             header              │
├─────────────────────────────────┤
│ ┌─────────────────────────────┐ │
│ │            card             │ │
│ └─────────────────────────────┘ │
│ ┌─────────────────────────────┐ │
│ │            card             │ │
│ └─────────────────────────────┘ │
└─────────────────────────────────┘
```

## Layout 3: Landing

Purpose: Marketing pages, storytelling, product showcases.

```css
┌─────────────────────────────────┐
│          header (nav)           │
├─────────────────────────────────┤
│                                 │
│         hero section            │
│    (full-width, split, etc)     │
│                                 │
├─────────────────────────────────┤
│                                 │
│       features section          │
│    (3-column, alternating)      │
│                                 │
├─────────────────────────────────┤
│                                 │
│     testimonials section        │
│         (carousel)              │
│                                 │
├─────────────────────────────────┤
│                                 │
│          cta section            │
│      (centered, button)         │
│                                 │
├─────────────────────────────────┤
│            footer               │
└─────────────────────────────────┘
```

## Layout 4: Application

Purpose: Complex tools with multiple panels (email client, IDE, admin).

### 4.1 Application - Desktop (≥1024px)

```css
┌─────────────────────────────────────┐
│              header                 │
├──────────┬──────────────┬───────────┤
│          │              │           │
│   left   │    main      │   right   │
│   nav    │   content    │   panel   │
│          │              │           │
│          │              │           │
├──────────┴──────────────┴───────────┤
│              footer?                │
└─────────────────────────────────────┘
```

### 4.2 Application - Tablet (640px - 1023px)

```css
┌─────────────────────────────────────┐
│              header                 │
├─────────────────────────────────────┤
│   ☰ menu button                     │
├──────────┬──────────────────────────┤
│          │                          │
│   left   │         main             │
│   nav    │       content            │
│  (drawer)│                          │
│          │                          │
└──────────┴──────────────────────────┘
```

### 4.3 Application - Mobile  (<640px)

```css
┌─────────────────────────────────────┐
│              header                 │
├─────────────────────────────────────┤
│                                     │
│                                     │
│              main                   │
│            content                  │
│                                     │
│                                     │
├─────────────────────────────────────┤
│  home   search   profile   settings │
└─────────────────────────────────────┘
```

## Layout 5: Blog

Purpose: Long-form reading, articles, posts.

### 5.1 Blog - Desktop (≥1024px)

```css
┌─────────────────────────────────┐
│            header               │
├─────────────────────────────────┤
│       featured image?           │
├────────────────┬────────────────┤
│                │                │
│    article     │   sidebar      │
│   (700px max)  │   (author,     │
│                │   related,     │
│                │     etc)       │
├────────────────┴────────────────┤
│       comments, related         │
├─────────────────────────────────┤
│            footer               │
└─────────────────────────────────┘
```

### 5.2 Blog - Mobile (<640px)

```css
┌─────────────────────────────────┐
│            header               │
├─────────────────────────────────┤
│       featured image?           │
├─────────────────────────────────┤
│                                 │
│           article               │
│         (full width)            │
│                                 │
├─────────────────────────────────┤
│           comments              │
├─────────────────────────────────┤
│            footer               │
└─────────────────────────────────┘
```

### Example Page Blog

```html
<body class="blog-layout">
  <header class="blog-layout__header">...</header>
  
  <img class="blog-layout__featured-image" src="post.jpg">
  
  <div class="blog-layout__content">
    <article class="blog-layout__article">
      <h1>Post Title</h1>
      <p>Content...</p>
    </article>
    
    <aside class="blog-layout__sidebar">
      <div class="author-card">...</div>
      <div class="related-posts">...</div>
    </aside>
  </div>
  
  <section class="blog-layout__comments">
    <h2>Comments</h2>
    ...
  </section>
  
  <footer class="blog-layout__footer">...</footer>
</body>
```

## Layout 6: Product

Purpose: E-commerce product pages.

### 6.1 Product - Desktop (≥1024px)

```css
┌─────────────────────────────────────┐
│               header                │
├─────────────────────────────────────┤
│ ┌───────────┐ ┌───────────────────┐ │
│ │           │ │                   │ │
│ │ gallery   │ │   product info    │ │
│ │           │ │   (title, price,  │ │
│ │           │ │    add to cart)   │ │
│ │           │ │                   │ │
│ └───────────┘ └───────────────────┘ │
├─────────────────────────────────────┤
│          product details            │
│    (specs, description, shipping)   │
├─────────────────────────────────────┤
│             reviews                 │
├─────────────────────────────────────┤
│          related products           │
├─────────────────────────────────────┤
│              footer                 │
└─────────────────────────────────────┘
```

### 6.2 Product - Mobile (<640px)

```css
┌─────────────────────────────────────┐
│               header                │
├─────────────────────────────────────┤
│ ┌─────────────────────────────────┐ │
│ │            gallery              │ │
│ └─────────────────────────────────┘ │
├─────────────────────────────────────┤
│           product info              │
│     (title, price, add to cart)     │
├─────────────────────────────────────┤
│          product details            │
├─────────────────────────────────────┤
│             reviews                 │
├─────────────────────────────────────┤
│         related products            │
├─────────────────────────────────────┤
│             footer                  │
└─────────────────────────────────────┘
```

### Example Page Product

```html
<body class="product-layout">
  <header class="product-layout__header">...</header>
  
  <main class="product-layout__main">
    <div class="product-layout__grid">
      <div class="product-layout__gallery">...</div>
      <div class="product-layout__info">...</div>
    </div>
    
    <section class="product-layout__section">
      <h2 class="product-layout__section-title">Details</h2>
      ...
    </section>
    
    <section class="product-layout__section">
      <h2 class="product-layout__section-title">Related Products</h2>
      <div class="product-layout__related">
        <div class="product-card">...</div>
        <div class="product-card">...</div>
        <div class="product-card">...</div>
        <div class="product-card">...</div>
      </div>
    </section>
  </main>
  
  <footer class="product-layout__footer">...</footer>
</body>
```

## Layout 7: Search/Listing

Purpose: Browse and filter content like products, articles, items.

### 7.1 Search - Desktop (≥1024px)

```css
┌─────────────────────────────────────┐
│              header                 │
├─────────────────────────────────────┤
│            search bar               │
├──────────┬──────────────────────────┤
│          │                          │
│ filters  │       results grid       │
│ (sidebar)│    (3-column cards)      │
│          │                          │
│          │                          │
└──────────┴──────────────────────────┘
```

### 7.2 Search - Mobile (<640px)

```css
┌─────────────────────────────────────┐
│               header                │
├─────────────────────────────────────┤
│             search bar              │
├─────────────────────────────────────┤
│    ☷ filter button (drawer)         │
├─────────────────────────────────────┤
│ ┌─────────────────────────────────┐ │
│ │            result card          │ │
│ └─────────────────────────────────┘ │
│ ┌─────────────────────────────────┐ │
│ │            result card          │ │
│ └─────────────────────────────────┘ │
│ ┌─────────────────────────────────┐ │
│ │            result card          │ │
│ └─────────────────────────────────┘ │
└─────────────────────────────────────┘
```

