# Web3 Portfolio Dashboard - Design System

## Overview
A modern, clean design system for managing Web3 portfolios with real-time data visualization and comprehensive analytics.

## Color Palette

### Primary Colors
- **Blue-600**: `#2563eb` - Primary brand color, buttons, links
- **Blue-500**: `#3b82f6` - Hover states, secondary actions
- **Blue-100**: `#dbeafe` - Background highlights, icons
- **Blue-50**: `#eff6ff` - Card backgrounds, subtle highlights

### Semantic Colors
- **Success Green**: `#16a34a` - Positive changes, success states
- **Warning Orange**: `#ea580c` - Warnings, attention states
- **Error Red**: `#dc2626` - Errors, negative changes
- **Info Blue**: `#0284c7` - Information, neutral states

### Neutral Colors
- **Gray-900**: `#111827` - Primary text
- **Gray-700**: `#374151` - Secondary text
- **Gray-500**: `#6b7280` - Tertiary text, labels
- **Gray-300**: `#d1d5db` - Borders, dividers
- **Gray-100**: `#f3f4f6` - Backgrounds
- **Gray-50**: `#f9fafb` - Page backgrounds

### Network-Specific Colors
- **Ethereum**: `#627eea` - Blue-purple
- **Polygon**: `#8247e5` - Purple
- **BSC**: `#f3ba2f` - Yellow
- **Arbitrum**: `#28a0f0` - Light blue

## Typography

### Font Family
- **Primary**: Inter (Google Fonts)
- **Monospace**: JetBrains Mono (for addresses, numbers)

### Font Weights
- **Regular**: 400 - Body text
- **Medium**: 500 - Labels, secondary text
- **Semibold**: 600 - Headings, important text
- **Bold**: 700 - Page titles, emphasis

### Font Sizes
- **xs**: 0.75rem (12px) - Captions, metadata
- **sm**: 0.875rem (14px) - Body text, labels
- **base**: 1rem (16px) - Default text
- **lg**: 1.125rem (18px) - Subheadings
- **xl**: 1.25rem (20px) - Section headings
- **2xl**: 1.5rem (24px) - Page titles
- **3xl**: 1.875rem (30px) - Hero text

## Component Library

### Buttons

#### Primary Button
```html
<button class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg font-medium transition-colors">
  Primary Action
</button>
```

#### Secondary Button
```html
<button class="bg-gray-200 hover:bg-gray-300 text-gray-700 px-4 py-2 rounded-lg font-medium transition-colors">
  Secondary Action
</button>
```

#### Icon Button
```html
<button class="p-2 text-gray-500 hover:text-gray-700 hover:bg-gray-100 rounded-lg transition-colors">
  <svg class="w-5 h-5">...</svg>
</button>
```

### Cards

#### Standard Card
```html
<div class="bg-white rounded-lg shadow p-6">
  <h3 class="text-lg font-semibold text-gray-900 mb-4">Card Title</h3>
  <p class="text-gray-600">Card content goes here</p>
</div>
```

#### Interactive Card
```html
<div class="bg-white rounded-lg shadow p-6 hover:shadow-lg transition-shadow cursor-pointer">
  <h3 class="text-lg font-semibold text-gray-900 mb-4">Interactive Card</h3>
  <p class="text-gray-600">Hover to see shadow effect</p>
</div>
```

#### Stats Card
```html
<div class="bg-white rounded-lg shadow p-6">
  <div class="flex items-center">
    <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center">
      <svg class="w-6 h-6 text-blue-600">...</svg>
    </div>
    <div class="ml-4">
      <p class="text-sm font-medium text-gray-500">Label</p>
      <p class="text-2xl font-bold text-gray-900">Value</p>
    </div>
  </div>
</div>
```

### Navigation

#### Sidebar Navigation
```html
<nav class="p-4 space-y-2">
  <a href="#" class="flex items-center px-4 py-3 rounded-lg text-blue-600 bg-blue-50 font-medium">
    <svg class="w-5 h-5 mr-3">...</svg>
    Active Item
  </a>
  <a href="#" class="flex items-center px-4 py-3 rounded-lg text-gray-700 hover:text-blue-600 font-medium">
    <svg class="w-5 h-5 mr-3">...</svg>
    Inactive Item
  </a>
</nav>
```

#### Mobile Navigation
```html
<div class="flex justify-around py-2">
  <a href="#" class="flex flex-col items-center p-2 text-blue-600">
    <svg class="w-6 h-6">...</svg>
    <span class="text-xs mt-1">Label</span>
  </a>
</div>
```

### Forms

#### Input Field
```html
<div class="space-y-2">
  <label class="block text-sm font-medium text-gray-700">Label</label>
  <input type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
</div>
```

#### Select Dropdown
```html
<select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
  <option>Option 1</option>
  <option>Option 2</option>
</select>
```

#### Toggle Switch
```html
<button class="bg-blue-600 relative inline-flex h-6 w-11 items-center rounded-full">
  <span class="inline-block h-4 w-4 transform rounded-full bg-white shadow translate-x-6"></span>
</button>
```

### Data Visualization

#### Progress Bar
```html
<div class="w-full bg-gray-200 rounded-full h-2">
  <div class="bg-blue-600 h-2 rounded-full" style="width: 45%"></div>
</div>
```

#### Network Allocation Chart
```html
<div class="space-y-4">
  <div class="flex items-center justify-between">
    <div class="flex items-center">
      <div class="w-3 h-3 rounded-full bg-blue-600 mr-3"></div>
      <span class="text-sm font-medium">Network</span>
    </div>
    <span class="text-sm text-gray-500">45%</span>
  </div>
  <div class="w-full bg-gray-200 rounded-full h-2">
    <div class="bg-blue-600 h-2 rounded-full" style="width: 45%"></div>
  </div>
</div>
```

#### Token Card
```html
<div class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-lg p-4">
  <div class="flex items-center justify-between mb-2">
    <span class="text-sm font-medium text-blue-900">ETH</span>
    <span class="text-xs text-blue-700">Ethereum</span>
  </div>
  <div class="text-2xl font-bold text-blue-900 mb-1">$4,875.00</div>
  <div class="text-sm text-green-600">+1.2%</div>
  <div class="text-xs text-blue-700 mt-1">2.5 ETH</div>
</div>
```

## Layout Patterns

### Page Structure
```html
<div class="min-h-screen flex">
  <!-- Sidebar -->
  <aside class="w-64 bg-white shadow-lg hidden md:block">
    <!-- Navigation content -->
  </aside>
  
  <!-- Main Content -->
  <div class="flex-1 flex flex-col">
    <!-- Header -->
    <header class="bg-white shadow-sm border-b border-gray-200 px-6 py-4">
      <!-- Header content -->
    </header>
    
    <!-- Main Content Area -->
    <main class="flex-1 p-6 space-y-6">
      <!-- Page content -->
    </main>
  </div>
</div>
```

### Grid Layouts

#### 4-Column Stats Grid
```html
<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
  <!-- Stats cards -->
</div>
```

#### 2-Column Content Grid
```html
<div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
  <!-- Content sections -->
</div>
```

#### 3-Column Token Grid
```html
<div class="grid grid-cols-1 md:grid-cols-3 gap-6">
  <!-- Token cards -->
</div>
```

## Responsive Design

### Breakpoints
- **Mobile**: < 768px
- **Tablet**: 768px - 1024px
- **Desktop**: > 1024px

### Mobile Adaptations
- Collapsible sidebar
- Bottom navigation bar
- Stacked grid layouts
- Simplified card designs

### Tablet Adaptations
- Sidebar remains visible
- Responsive grid adjustments
- Touch-friendly button sizes

## Animation & Interactions

### Transitions
- **Default**: 0.2s ease
- **Hover**: 0.3s ease
- **Page transitions**: 0.35s easeInOut

### Hover Effects
```css
.card-hover {
  transition: all 0.3s ease;
}

.card-hover:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 25px rgba(0,0,0,0.1);
}
```

### Loading States
```html
<div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
```

## Accessibility

### Color Contrast
- All text meets WCAG AA standards (4.5:1 ratio)
- Interactive elements have sufficient contrast
- Color is not the only indicator of state

### Keyboard Navigation
- Full keyboard support
- Visible focus indicators
- Logical tab order

### Screen Reader Support
- Proper ARIA labels
- Semantic HTML structure
- Descriptive alt text for images

## Dark Mode Support

### Color Mapping
- Background: `gray-50` → `gray-900`
- Cards: `white` → `gray-950`
- Text: `gray-900` → `gray-100`
- Borders: `gray-200` → `gray-800`

### Implementation
```css
.dark {
  --bg-primary: #111827;
  --bg-secondary: #1f2937;
  --text-primary: #f9fafb;
  --text-secondary: #d1d5db;
}
```

## Icon System

### Icon Sizes
- **Small**: 16px (w-4 h-4)
- **Medium**: 20px (w-5 h-5)
- **Large**: 24px (w-6 h-6)
- **Extra Large**: 32px (w-8 h-8)

### Icon Colors
- **Primary**: `text-blue-600`
- **Secondary**: `text-gray-500`
- **Success**: `text-green-600`
- **Warning**: `text-orange-600`
- **Error**: `text-red-600`

## Spacing System

### Spacing Scale
- **xs**: 0.25rem (4px)
- **sm**: 0.5rem (8px)
- **md**: 1rem (16px)
- **lg**: 1.5rem (24px)
- **xl**: 2rem (32px)
- **2xl**: 3rem (48px)

### Common Spacing Patterns
- **Card padding**: `p-6` (24px)
- **Section spacing**: `space-y-6` (24px between items)
- **Grid gaps**: `gap-6` (24px)
- **Button padding**: `px-4 py-2` (16px horizontal, 8px vertical)

This design system provides a consistent, accessible, and modern foundation for the Web3 Portfolio Dashboard, ensuring a great user experience across all devices and use cases. 