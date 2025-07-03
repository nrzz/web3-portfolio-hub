# 🌿 Git Branching Strategy

This project follows a **Git Flow** branching model with two main branches for organized development and stable releases.

## 📋 Branch Structure

### 🏠 **Main Branch (`main`)**
- **Purpose**: Stable, production-ready code
- **Protection**: Protected branch (no direct pushes)
- **Deployment**: Automatically deployed to production
- **Merging**: Only via Pull Requests from `dev` branch

### 🚧 **Development Branch (`dev`)**  
- **Purpose**: Active development and integration
- **Source**: Cloned from `main` branch
- **Features**: All new features merge here first
- **Testing**: Continuous integration and testing
- **Merging**: Features merge via Pull Requests

## 🔄 Workflow Process

### 1. **Feature Development**
```bash
# Start from dev branch
git checkout dev
git pull origin dev

# Create feature branch
git checkout -b feature/your-feature-name

# Work on your feature
git add .
git commit -m "Add: new feature implementation"

# Push feature branch
git push origin feature/your-feature-name
```

### 2. **Feature Integration**
- Create **Pull Request** from `feature/xyz` → `dev`
- Code review and testing
- Merge to `dev` branch
- Delete feature branch

### 3. **Release Process**
```bash
# Create release branch from dev
git checkout dev
git checkout -b release/v1.0.0

# Final testing and bug fixes
git commit -m "Fix: release preparation"

# Merge to main via PR
# Create PR: release/v1.0.0 → main
```

### 4. **Hotfixes** (Emergency fixes to production)
```bash
# Create hotfix from main
git checkout main
git checkout -b hotfix/critical-bug-fix

# Fix the issue
git commit -m "Fix: critical production bug"

# Merge to both main and dev
```

## 🛡️ Branch Protection Rules

### **Main Branch**
- ✅ Require pull request reviews
- ✅ Require status checks to pass
- ✅ Require up-to-date branches
- ✅ Include administrators
- ❌ No direct pushes allowed

### **Dev Branch**
- ✅ Require pull request reviews (recommended)
- ✅ Require status checks to pass
- ✅ Allow administrators to bypass

## 📁 Branch Naming Conventions

### **Feature Branches**
- `feature/user-authentication`
- `feature/portfolio-analytics`
- `feature/forum-voting-system`

### **Bug Fix Branches**
- `bugfix/login-validation`
- `bugfix/responsive-layout`

### **Hotfix Branches**
- `hotfix/security-vulnerability`
- `hotfix/database-connection`

### **Release Branches**
- `release/v1.0.0`
- `release/v1.1.0`

## 🚀 Quick Commands

### **Switch to Dev Branch**
```bash
git checkout dev
git pull origin dev
```

### **Create New Feature**
```bash
git checkout dev
git checkout -b feature/new-feature
```

### **Update from Main**
```bash
git checkout dev
git pull origin main
```

### **Check Current Branch**
```bash
git branch
git status
```

## 📊 Current Branch Status

- **Main Branch**: `main` - Stable production code
- **Development Branch**: `dev` - Active development
- **Current Branch**: `dev` (ready for feature development)

## 🤝 Contribution Guidelines

1. **Always start from `dev`** for new features
2. **Create descriptive commit messages**
3. **Test your code** before creating PRs
4. **Keep PRs focused** on single features
5. **Update documentation** when needed
6. **Follow code review feedback**

---

**Happy coding! 🎉** 