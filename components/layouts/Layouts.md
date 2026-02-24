# Layout System Documentation

## Core Philosophy

The folder provides a library of opinionated, named layouts. Each layout serves a specific content type and knows how to rearrange itself at different screen sizes.

## Layout 1: Documentation

Purpose: Content with navigation/supporting info (docs, tutorials, reference).  
Documentation HTML file provides an example that includes debug info.

### 1.1. Documentation - Desktop (≥1024px)

```go
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

```go
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

```go
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

```go
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

```go
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
Landing HTML file provides an example that includes debug info with the following structure:

- Sticky Header - Logo + navigation
- Hero Section - Split layout with CTA and image placeholder
- Features Section - 3-column grid of feature cards (stacks on mobile)
- Testimonials - 3-column with author avatars
- Logo Cloud - 4-column company logos (2-column on tablet, 1 on mobile)
- Pricing - 3-column cards with "popular" highlighted
- CTA Section - Full-width colored section
- Footer - 4-column links

### Structure

```go
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
HTML files provide an example that include debug info:

04_application.html                 (Main application dashboard)
04_application_analytics.html       (Analytics page)
04_application_projects.html        (Projects list page)
04_application_project.html         (Single project view)
04_application_tasks.html           (Tasks list page)
04_application_task.html            (Single task view)
04_application_team.html            (Team members list)
04_application_team_member.html     (Single team member view)
04_application_reports.html         (Reports list page)
04_application_report.html          (Single report view/definition)
04_application_report_run.html      (Run report page)
04_application_report_target.html   (Report target configuration)
04_application_reports_comparison.html (Compare reports page)

### 4.1 Application - Desktop (≥1024px)

```go
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

```go
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

```go
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
Blog HTML file provides an example that includes debug info.

### 5.1 Blog - Desktop (≥1024px)

```go
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

```go
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

## Layout 6: Product

Purpose: E-commerce product pages.  
Product HTML file provides an example that includes debug info with the following structure:

- Header - Logo, search, navigation, cart
- Product Grid - Gallery (left) + Info (right)
- Product Details - Tabs with description and specs
- Reviews Section - Rating summary and individual reviews
- Related Products - 4-column grid of related items
- Footer - 4-column links

### 6.1 Product - Desktop (≥1024px)

```go
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

```go
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

## Layout 7: Search/Listing

Purpose: Browse and filter content like products, articles, items.  
Search products HTML file provides an example that includes debug info with the following features:

- Sticky Filters on desktop/tablet
- Mobile-friendly filter toggle
- Product cards with image, category, rating, price, badges
- Sort dropdown for results
- Pagination at bottom
- Clear filter option
- Price range inputs
- Checkbox filters for categories, brands, features

### 7.1 Search - Desktop (≥1024px)

```go
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

```go
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
