# Blog details page

## Page Layout

### 1. Container

The container is the horizontal boundary of the page. It ensures that content is readable on all screen sizes and never stretches too wide.

Purpose:  
a. Keeps text at a readable width  
b. Centers the page content  
c. Provides consistent horizontal padding  
d. Adapts fluidly to screen size

Behavior:  
a. Full‑width background, but content stays centered  
b. Fluid width on mobile (100% minus padding)  
c. Max‑width on tablet and desktop (e.g., 1200px)  
d. Always includes left/right padding so content never touches the edges

### 2. Sections

Sections are the vertical building blocks of every page. They define how content is grouped, spaced, and visually separated from other parts of the layout.

Purpose:
a. Break the page into clear, readable chunks  
b. Control vertical rhythm  
c. Allow each part of the page to have its own background, spacing, and layout  
d. Make the page scannable and predictable

Behavior:
a. Sections stack top‑to‑bottom  
b. Each section contains its own container  
c. Each section has consistent vertical spacing (small / medium / large)  
d. Sections may define their own background (light, dark, accent)

### 3. Grid

The grid defines how content splits into columns at different screen sizes. It is the core mechanism that makes the layout responsive and readable.  
The grid ensures that every section — hero, content, sidebar, cards, footer — behaves consistently across breakpoints.

Purpose:  
a. Control horizontal structure  
b. Adapt column count to screen size  
c. Keep content balanced and predictable  
d. Support both simple and complex layouts

Behavior
a. Mobile: always 1 column (stacked, full‑width)  
b. Tablet: 1–2 columns depending on the section  
c. Desktop: 2–3 columns depending on the design  
d. Columns always live inside the container. Columns have consistent gaps defined by the spacing scale

### 4. Spacing

Spacing defines the vertical and horizontal rhythm of the layout. It ensures that all sections, components, and text blocks feel consistent and balanced across the entire site.  
A spacing system makes the layout feel intentional and cohesive.
It reduces design drift, simplifies CSS, and ensures that every widget aligns with the overall rhythm of the site.

Purpose:  
a. Establish a predictable visual rhythm  
b. Prevent components from feeling cramped or uneven  
c. Create consistent breathing room across all pages  
d. Support readability on both small and large screens

Behavior:  
a. Spacing follows a small / medium / large scale  
b. Vertical spacing between sections uses the same scale  
c. Horizontal spacing inside the container is consistent at all breakpoints  
d. Spacing increases slightly on larger screens to match wider layouts  
e. Components never define arbitrary spacing; they use the shared scale

### 5. Typography Flow

Typography flow defines how text scales and reflows across different screen sizes. It ensures that reading remains comfortable on mobile, tablet, and desktop without requiring separate designs.  
Typography flow ensures that every page — blog posts, landing pages, service pages — feels readable and balanced. It ties together spacing, container width, and grid behavior into a coherent reading experience.

Purpose:  
a. Maintain readability across all breakpoints  
b. Establish a clear hierarchy between headings, subheadings, and body text  
c. Prevent text from feeling cramped on mobile or oversized on desktop  
d. Support consistent rhythm with the spacing system

Behavior:  
a. Font sizes follow a small / medium / large scale that increases on larger screens  
b. Line‑height adjusts to maintain readability as text scales  
c. Headings scale proportionally so hierarchy remains clear  
d. Paragraph width is constrained by the container to avoid long line lengths  
e. Typography changes are minimal and predictable, not dramatic shifts

## Header layout

The header is the global navigation frame of the site. It appears on every page and provides identity, orientation, and access to primary actions.  
The header is the first structural component users interact with.
A stable, predictable header reinforces branding and provides a clear entry point into the site’s navigation system.

Purpose:  
a. Establish brand identity at the top of the page  
b. Provide consistent navigation across the site  
c. Anchor the layout with a stable, predictable structure  
d. Remain usable on all screen sizes

Behavior:  
a. Contains the logo, primary navigation, and optional actions (search, account, etc.)  
b. Uses the container to align content horizontally  
c. On mobile, navigation collapses into a simplified or stacked layout  
d. On desktop, navigation displays inline with full spacing  
e. Maintains consistent height, spacing, and typography across pages  

## Breadcrumb / Hero section

The breadcrumb and hero area provide orientation and context at the top of the page. They tell the reader where they are and what the page is about before any content begins.  
This section sets expectations and reduces cognitive load.
It helps users understand the structure of the site and the purpose of the page before they begin reading.

Purpose:  
a. Show the user’s position within the site hierarchy  
b. Present the page title or article title prominently  
c. Provide immediate context for the content that follows  
d. Establish visual hierarchy at the top of the page

Behavior:  
a. Breadcrumbs appear above the title on desktop and may collapse or simplify on mobile  
b. The hero section contains the main title, optional subtitle, and optional supporting elements  
c. The section uses the container for alignment and spacing  
d. Typography scales with breakpoints to maintain clarity  
e. Background may differ from the rest of the page to create visual separation

## Content column vs Sidebar grid

The content–sidebar grid defines the primary reading structure of the blog details page. It balances long‑form content with supporting information while adapting cleanly across breakpoints.  
This grid is the backbone of the blog details page.
It ensures that long‑form content remains readable while still allowing secondary information to be visible without overwhelming the main text.

Purpose:  
a. Provide a clear reading column for the article body  
b. Offer space for secondary elements such as related posts, categories, or calls to action  
c. Maintain readability and hierarchy across all screen sizes

Behavior:  
a. Mobile: single‑column layout; sidebar content stacks below the article  
b. Tablet: optional two‑column layout depending on available width  
c. Desktop: stable two‑column layout with a dominant content column and a narrower sidebar  
d. Column widths remain consistent across pages to reinforce structure  
e. Column spacing follows the global spacing scale

## Article metadata block

The article metadata block presents key contextual information about the blog post. It helps readers understand the author, timing, and categorization before entering the main content.  
Metadata sets expectations and orients the reader.
A clear, consistent metadata block improves usability and reinforces the structure of the blog details page.

Purpose:  
a. Provide essential context such as author, publication date, reading time, and categories  
b. Establish trust and credibility  
c. Help readers quickly assess relevance

Behavior:  
a. Appears directly below the hero/title section  
b. Uses a compact, consistent layout that aligns with the content column  
c. Metadata items follow a predictable order and spacing  
d. On mobile, items stack vertically for clarity  
e. On desktop, items may appear inline or grouped for efficient scanning

## Article body structure

The article body structure defines how long‑form content is presented inside the main reading column. It ensures that headings, paragraphs, images, and embedded elements follow a consistent, readable pattern.  
A consistent article body structure makes long‑form content easy to read and visually coherent. It ensures that every blog post feels part of the same system, regardless of author or topic.

Purpose:  
a. Provide a clean, predictable reading experience  
b. Maintain consistent hierarchy across all article types  
c. Support rich content such as images, quotes, and code blocks  
d. Ensure readability on both mobile and desktop

Behavior:  
a. The body lives entirely inside the content column  
b. Headings follow a clear hierarchy (H1 → H2 → H3)  
c. Paragraph width is constrained by the container for optimal line length  
d. Images and media scale fluidly and never overflow the column  
e. Lists, quotes, and code blocks follow consistent spacing rules  
f. Typography and spacing adjust slightly across breakpoints to maintain readability

## Footer layout

The footer is the closing structural element of the page. It provides navigation, legal information, and secondary content while maintaining visual balance at the bottom of the layout.  
The footer anchors the page and provides closure.
A well‑structured footer improves navigation, reinforces trust, and ensures that every page ends with a consistent, predictable layout.

Purpose:  
a. Offer global links such as About, Contact, Privacy, and Terms  
b. Reinforce branding at the end of the page  
c. Provide a clear visual endpoint to the reading experience  
d. Maintain consistency across all pages

Behavior:  
a. Uses the container for horizontal alignment  
b. Organizes content into stacked sections on mobile and columns on desktop  
c. Includes consistent spacing above and below to separate it from the article body  
d. Typography and spacing follow the global scales  
e. May include a secondary background to visually distinguish it from the main content
