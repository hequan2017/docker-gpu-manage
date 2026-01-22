/**
 * 天启算力管理平台 - 官网交互脚本
 * 提供导航、动画、标签页切换等功能
 */

// ===== 等待 DOM 加载完成 =====
document.addEventListener('DOMContentLoaded', function() {

    // ===== 导航栏滚动效果 =====
    const navbar = document.querySelector('.navbar');
    let lastScrollY = window.scrollY;

    window.addEventListener('scroll', function() {
        const currentScrollY = window.scrollY;

        // 添加/移除滚动样式
        if (currentScrollY > 50) {
            navbar.classList.add('scrolled');
        } else {
            navbar.classList.remove('scrolled');
        }

        lastScrollY = currentScrollY;
    });

    // ===== 移动端菜单切换 =====
    const navToggle = document.querySelector('.nav-toggle');
    const navMenu = document.querySelector('.nav-menu');

    if (navToggle) {
        navToggle.addEventListener('click', function() {
            navMenu.classList.toggle('active');
            this.classList.toggle('active');
        });
    }

    // 点击导航链接后关闭移动端菜单
    const navLinks = document.querySelectorAll('.nav-link');
    navLinks.forEach(link => {
        link.addEventListener('click', function() {
            if (window.innerWidth <= 768) {
                navMenu.classList.remove('active');
                navToggle.classList.remove('active');
            }
        });
    });

    // ===== 平滑滚动到锚点 =====
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function(e) {
            const href = this.getAttribute('href');

            // 跳过空链接和外部链接
            if (href === '#' || href.startsWith('http')) {
                return;
            }

            e.preventDefault();
            const target = document.querySelector(href);

            if (target) {
                const navbarHeight = navbar.offsetHeight;
                const targetPosition = target.getBoundingClientRect().top + window.scrollY - navbarHeight;

                window.scrollTo({
                    top: targetPosition,
                    behavior: 'smooth'
                });
            }
        });
    });

    // ===== 导航链接高亮 =====
    const sections = document.querySelectorAll('section[id]');

    function updateActiveLink() {
        const scrollY = window.scrollY;

        sections.forEach(section => {
            const sectionTop = section.offsetTop - navbar.offsetHeight - 100;
            const sectionHeight = section.offsetHeight;
            const sectionId = section.getAttribute('id');

            if (scrollY > sectionTop && scrollY <= sectionTop + sectionHeight) {
                navLinks.forEach(link => {
                    link.classList.remove('active');
                    if (link.getAttribute('href') === `#${sectionId}`) {
                        link.classList.add('active');
                    }
                });
            }
        });
    }

    window.addEventListener('scroll', updateActiveLink);
    updateActiveLink(); // 初始化

    // ===== 数字动画效果 =====
    function animateNumbers() {
        const statNumbers = document.querySelectorAll('.stat-number');

        statNumbers.forEach(stat => {
            const target = parseInt(stat.getAttribute('data-target'));
            const duration = 2000; // 动画持续时间（毫秒）
            const step = target / (duration / 16); // 60fps
            let current = 0;

            const updateNumber = () => {
                current += step;
                if (current < target) {
                    stat.textContent = Math.floor(current);
                    requestAnimationFrame(updateNumber);
                } else {
                    stat.textContent = target;
                }
            };

            // 使用 Intersection Observer 检测元素是否在视口中
            const observer = new IntersectionObserver((entries) => {
                entries.forEach(entry => {
                    if (entry.isIntersecting && !stat.classList.contains('animated')) {
                        stat.classList.add('animated');
                        updateNumber();
                    }
                });
            }, { threshold: 0.5 });

            observer.observe(stat);
        });
    }

    animateNumbers();

    // ===== 标签页切换 =====
    const tabButtons = document.querySelectorAll('.tab-btn');
    const tabPanes = document.querySelectorAll('.tab-pane');

    tabButtons.forEach(button => {
        button.addEventListener('click', function() {
            const tabId = this.getAttribute('data-tab');

            // 移除所有活动状态
            tabButtons.forEach(btn => btn.classList.remove('active'));
            tabPanes.forEach(pane => pane.classList.remove('active'));

            // 添加当前活动状态
            this.classList.add('active');
            const targetPane = document.getElementById(tabId);
            if (targetPane) {
                targetPane.classList.add('active');
            }
        });
    });

    // ===== 代码复制功能 =====
    const copyButtons = document.querySelectorAll('.copy-btn');

    copyButtons.forEach(button => {
        button.addEventListener('click', function() {
            const codeId = this.getAttribute('data-code');
            const codeElement = document.getElementById(codeId);

            if (codeElement) {
                const code = codeElement.textContent;

                // 使用 Clipboard API 复制
                navigator.clipboard.writeText(code).then(() => {
                    // 显示复制成功状态
                    const originalHTML = this.innerHTML;
                    this.innerHTML = '<i class="fas fa-check"></i><span>已复制</span>';
                    this.classList.add('copied');

                    // 2秒后恢复原始状态
                    setTimeout(() => {
                        this.innerHTML = originalHTML;
                        this.classList.remove('copied');
                    }, 2000);
                }).catch(err => {
                    // 降级方案：创建临时文本框
                    const textArea = document.createElement('textarea');
                    textArea.value = code;
                    textArea.style.position = 'fixed';
                    textArea.style.opacity = '0';
                    document.body.appendChild(textArea);
                    textArea.select();

                    try {
                        document.execCommand('copy');
                        const originalHTML = this.innerHTML;
                        this.innerHTML = '<i class="fas fa-check"></i><span>已复制</span>';
                        this.classList.add('copied');

                        setTimeout(() => {
                            this.innerHTML = originalHTML;
                            this.classList.remove('copied');
                        }, 2000);
                    } catch (e) {
                        console.error('复制失败:', e);
                    }

                    document.body.removeChild(textArea);
                });
            }
        });
    });

    // ===== 滚动动画 =====
    const observerOptions = {
        threshold: 0.1,
        rootMargin: '0px 0px -100px 0px'
    };

    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.style.animationPlayState = 'running';
                observer.unobserve(entry.target);
            }
        });
    }, observerOptions);

    // 观察所有需要动画的元素
    document.querySelectorAll('[data-aos]').forEach(el => {
        el.style.animationPlayState = 'paused';
        observer.observe(el);
    });

    // ===== 功能卡片悬浮效果增强 =====
    const featureCards = document.querySelectorAll('.feature-card');

    featureCards.forEach(card => {
        card.addEventListener('mousemove', function(e) {
            const rect = card.getBoundingClientRect();
            const x = e.clientX - rect.left;
            const y = e.clientY - rect.top;

            const centerX = rect.width / 2;
            const centerY = rect.height / 2;

            const rotateX = (y - centerY) / 20;
            const rotateY = (centerX - x) / 20;

            card.style.transform = `perspective(1000px) rotateX(${rotateX}deg) rotateY(${rotateY}deg) translateY(-5px)`;
        });

        card.addEventListener('mouseleave', function() {
            card.style.transform = 'perspective(1000px) rotateX(0) rotateY(0) translateY(0)';
        });
    });

    // ===== 场景卡片悬浮效果增强 =====
    const scenarioCards = document.querySelectorAll('.scenario-card');

    scenarioCards.forEach(card => {
        card.addEventListener('mouseenter', function() {
            this.style.borderColor = 'rgba(0, 240, 255, 0.5)';
        });

        card.addEventListener('mouseleave', function() {
            this.style.borderColor = 'rgba(0, 240, 255, 0.2)';
        });
    });

    // ===== 技术图标动画 =====
    const techItems = document.querySelectorAll('.tech-item');

    techItems.forEach(item => {
        item.addEventListener('mouseenter', function() {
            const icon = this.querySelector('.tech-icon');
            if (icon) {
                icon.style.transform = 'scale(1.1) rotate(5deg)';
            }
        });

        item.addEventListener('mouseleave', function() {
            const icon = this.querySelector('.tech-icon');
            if (icon) {
                icon.style.transform = 'scale(1) rotate(0deg)';
            }
        });
    });

    // ===== 页面加载进度指示器 =====
    const createLoadingIndicator = () => {
        const loader = document.createElement('div');
        loader.className = 'page-loader';
        loader.innerHTML = `
            <div class="loader-spinner"></div>
            <div class="loader-text">加载中...</div>
        `;
        document.body.appendChild(loader);

        return loader;
    };

    // 在页面加载时显示加载指示器
    const loader = createLoadingIndicator();

    window.addEventListener('load', function() {
        setTimeout(() => {
            loader.classList.add('loaded');
            setTimeout(() => {
                document.body.removeChild(loader);
            }, 500);
        }, 500);
    });

    // ===== 添加滚动到顶部功能 =====
    const createScrollToTop = () => {
        const scrollBtn = document.createElement('button');
        scrollBtn.className = 'scroll-to-top';
        scrollBtn.innerHTML = '<i class="fas fa-arrow-up"></i>';
        scrollBtn.setAttribute('aria-label', '滚动到顶部');
        document.body.appendChild(scrollBtn);

        // 显示/隐藏按钮
        window.addEventListener('scroll', function() {
            if (window.scrollY > 500) {
                scrollBtn.classList.add('visible');
            } else {
                scrollBtn.classList.remove('visible');
            }
        });

        // 点击滚动到顶部
        scrollBtn.addEventListener('click', function() {
            window.scrollTo({
                top: 0,
                behavior: 'smooth'
            });
        });
    };

    createScrollToTop();

    // ===== 键盘快捷键支持 =====
    document.addEventListener('keydown', function(e) {
        // ESC 键关闭移动端菜单
        if (e.key === 'Escape') {
            navMenu.classList.remove('active');
            navToggle.classList.remove('active');
        }

        // Ctrl/Cmd + K 打开搜索（预留功能）
        if ((e.ctrlKey || e.metaKey) && e.key === 'k') {
            e.preventDefault();
            // TODO: 实现搜索功能
            console.log('搜索功能待实现');
        }
    });

    // ===== 性能优化：节流函数 =====
    function throttle(func, delay) {
        let timeoutId;
        let lastExecTime = 0;

        return function(...args) {
            const currentTime = Date.now();
            const timeSinceLastExec = currentTime - lastExecTime;

            clearTimeout(timeoutId);

            if (timeSinceLastExec > delay) {
                func.apply(this, args);
                lastExecTime = currentTime;
            } else {
                timeoutId = setTimeout(() => {
                    func.apply(this, args);
                    lastExecTime = Date.now();
                }, delay - timeSinceLastExec);
            }
        };
    }

    // 使用节流优化滚动事件
    const handleScroll = throttle(function() {
        // 滚动相关的处理逻辑
        updateActiveLink();
    }, 100);

    window.addEventListener('scroll', handleScroll);

    // ===== 控制彩蛋（可选） =====
    const KonamiCode = ['ArrowUp', 'ArrowUp', 'ArrowDown', 'ArrowDown', 'ArrowLeft', 'ArrowRight', 'ArrowLeft', 'ArrowRight', 'b', 'a'];
    let KonamiIndex = 0;

    document.addEventListener('keydown', function(e) {
        if (e.key === KonamiCode[KonamiIndex]) {
            KonamiIndex++;
            if (KonamiIndex === KonamiCode.length) {
                activateEasterEgg();
                KonamiIndex = 0;
            }
        } else {
            KonamiIndex = 0;
        }
    });

    function activateEasterEgg() {
        // 彩蛋效果：粒子爆发
        document.body.style.animation = 'rainbow 2s ease';
        setTimeout(() => {
            document.body.style.animation = '';
        }, 2000);
    }

    // ===== 调试信息（开发模式） =====
    const isDevelopment = window.location.hostname === 'localhost' || window.location.hostname === '127.0.0.1';

    if (isDevelopment) {
        console.log('%c天启算力管理平台', 'color: #00f0ff; font-size: 24px; font-weight: bold;');
        console.log('%c官网版本: 1.0.0', 'color: #7b2cbf; font-size: 14px;');
        console.log('%c技术栈: HTML5 + CSS3 + Vanilla JavaScript', 'color: #a0aec0; font-size: 12px;');
    }

    // ===== 页面可见性变化处理 =====
    document.addEventListener('visibilitychange', function() {
        if (document.hidden) {
            // 页面隐藏时暂停动画
            document.body.classList.add('page-hidden');
        } else {
            // 页面显示时恢复动画
            document.body.classList.remove('page-hidden');
        }
    });

    // ===== 初始化完成提示 =====
    console.log('%c✓ 官网初始化完成', 'color: #10b981; font-size: 14px;');
});

// ===== CSS 动画扩展（通过 JS 动态添加） =====
const additionalStyles = document.createElement('style');
additionalStyles.textContent = `
    /* 页面加载器 */
    .page-loader {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background: rgba(10, 14, 39, 0.95);
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        z-index: 9999;
        transition: opacity 0.5s ease;
    }

    .page-loader.loaded {
        opacity: 0;
    }

    .loader-spinner {
        width: 50px;
        height: 50px;
        border: 3px solid rgba(0, 240, 255, 0.2);
        border-top-color: #00f0ff;
        border-radius: 50%;
        animation: spin 1s linear infinite;
    }

    @keyframes spin {
        to { transform: rotate(360deg); }
    }

    .loader-text {
        margin-top: 1rem;
        color: #00f0ff;
        font-size: 14px;
        font-weight: 500;
    }

    /* 滚动到顶部按钮 */
    .scroll-to-top {
        position: fixed;
        bottom: 30px;
        right: 30px;
        width: 50px;
        height: 50px;
        background: linear-gradient(135deg, var(--primary-color), var(--primary-dark));
        border: none;
        border-radius: 50%;
        color: var(--bg-primary);
        font-size: 20px;
        cursor: pointer;
        opacity: 0;
        visibility: hidden;
        transition: all 0.3s ease;
        box-shadow: 0 4px 12px rgba(0, 240, 255, 0.3);
        z-index: 1000;
    }

    .scroll-to-top.visible {
        opacity: 1;
        visibility: visible;
    }

    .scroll-to-top:hover {
        transform: translateY(-5px);
        box-shadow: 0 6px 20px rgba(0, 240, 255, 0.5);
    }

    /* 页面隐藏时暂停动画 */
    .page-hidden *,
    .page-hidden *::before,
    .page-hidden *::after {
        animation-play-state: paused !important;
    }

    /* 彩蛋彩虹效果 */
    @keyframes rainbow {
        0% { filter: hue-rotate(0deg); }
        100% { filter: hue-rotate(360deg); }
    }

    /* 移动端菜单动画 */
    .nav-toggle.active span:nth-child(1) {
        transform: rotate(45deg) translate(5px, 5px);
    }

    .nav-toggle.active span:nth-child(2) {
        opacity: 0;
    }

    .nav-toggle.active span:nth-child(3) {
        transform: rotate(-45deg) translate(7px, -6px);
    }
`;

document.head.appendChild(additionalStyles);
