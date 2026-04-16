import type { CSSProperties, JSX } from "react"
import { Link } from "react-router-dom"

const statsContainerStyle: CSSProperties = {
    flexWrap: "wrap",
    marginBottom: "40px",
}

const statsBoxStyle: CSSProperties = {
    flex: 1,
    minWidth: "200px",
}

const mentorPlaceholderStyle: CSSProperties = {
    backgroundColor: "#e0e0e0",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    color: "#666",
}

const mentorImageStyle = (imageUrl: string): CSSProperties => ({
    backgroundImage: `url('${imageUrl}')`,
})

const MentorsPage = (): JSX.Element => {
    return (
        <main>
            <section className="section section--mentors">
                <div className="container">
                    <div className="section__header">
                        <div>
                            <h2 className="section__title">Наставники</h2>
                            <p className="section__subtitle">Профессиональные менторы по направлениям: Python, C#, Go, JavaScript</p>
                        </div>
                        <button className="btn btn--sm btn--outline">Сбросить</button>
                    </div>

                    <div className="stats__container" style={statsContainerStyle}>
                        <div className="stats__box" style={statsBoxStyle}>
                            <div className="stats__label">Менторов доступно</div>
                            <div className="stats__value">12</div>
                            <div className="stats__unit">👥</div>
                        </div>
                        <div className="stats__box" style={statsBoxStyle}>
                            <div className="stats__label">Средний рейтинг</div>
                            <div className="stats__value">4.8</div>
                            <div className="stats__unit">⭐</div>
                        </div>
                        <div className="stats__box" style={statsBoxStyle}>
                            <div className="stats__label">Консультаций проведено</div>
                            <div className="stats__value">350+</div>
                            <div className="stats__unit">🎓</div>
                        </div>
                    </div>

                    <div className="mentors-grid">
                        <div className="mentor-card">
                            <div className="mentor-card__top">
                                <span className="tag">Python</span>
                            </div>
                            <div className="mentor-card__image"
                                style={mentorImageStyle("https://images.unsplash.com/photo-1531891437562-4301cf35b7e4?ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80")}>
                            </div>
                            <div className="mentor-card__info">
                                <h3 className="mentor-card__title">Алексей Смирнов</h3>
                                <p className="mentor-card__role">Backend, Django, FastAPI, архитектура сервисов.</p>
                            </div>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Записаться</Link>
                            </div>
                        </div>
                        <div className="mentor-card">
                            <div className="mentor-card__top">
                                <span className="tag">Python</span>
                            </div>
                            <div className="mentor-card__image"
                                style={mentorImageStyle("https://images.unsplash.com/photo-1544716278-ca5e3f4abd8c?ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80")}>
                            </div>
                            <div className="mentor-card__info">
                                <h3 className="mentor-card__title">Мария Орлова</h3>
                                <p className="mentor-card__role">Highload, Kafka, PostgreSQL, DevOps-практики.</p>
                            </div>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Записаться</Link>
                            </div>
                        </div>
                        <div className="mentor-card">
                            <div className="mentor-card__top">
                                <span className="tag">C#</span>
                            </div>
                            <div className="mentor-card__image"
                                style={mentorImageStyle("https://images.unsplash.com/photo-1560250097-0b93528c311a?ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80")}>
                            </div>
                            <div className="mentor-card__info">
                                <h3 className="mentor-card__title">Иван Кузнецов</h3>
                                <p className="mentor-card__role">.NET, ASP.NET Core, микросервисы, Azure.</p>
                            </div>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Записаться</Link>
                            </div>
                        </div>
                        <div className="mentor-card">
                            <div className="mentor-card__top">
                                <span className="tag">C#</span>
                            </div>
                            <div className="mentor-card__image"
                                style={mentorPlaceholderStyle}>
                                Нет фото</div>
                            <div className="mentor-card__info">
                                <h3 className="mentor-card__title">Сергей Лебедев</h3>
                                <p className="mentor-card__role">DDD, CQRS, gRPC, оптимизация EF Core.</p>
                            </div>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Записаться</Link>
                            </div>
                        </div>
                        <div className="mentor-card">
                            <div className="mentor-card__top">
                                <span className="tag">Go</span>
                            </div>
                            <div className="mentor-card__image"
                                style={mentorPlaceholderStyle}>
                                Нет фото</div>
                            <div className="mentor-card__info">
                                <h3 className="mentor-card__title">Андрей Козлов</h3>
                                <p className="mentor-card__role">Go, gRPC, Kubernetes, наблюдаемость.</p>
                            </div>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Записаться</Link>
                            </div>
                        </div>
                        <div className="mentor-card">
                            <div className="mentor-card__top">
                                <span className="tag">Go</span>
                            </div>
                            <div className="mentor-card__image"
                                style={mentorPlaceholderStyle}>
                                Нет фото</div>
                            <div className="mentor-card__info">
                                <h3 className="mentor-card__title">Олег Морозов</h3>
                                <p className="mentor-card__role">Производительность, устойчивость, очереди.</p>
                            </div>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Записаться</Link>
                            </div>
                        </div>
                        <div className="mentor-card">
                            <div className="mentor-card__top">
                                <span className="tag">JS</span>
                            </div>
                            <div className="mentor-card__image"
                                style={mentorImageStyle("https://images.unsplash.com/photo-1517694712202-14dd9538aa97?ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80")}>
                            </div>
                            <div className="mentor-card__info">
                                <h3 className="mentor-card__title">Светлана Яковлева</h3>
                                <p className="mentor-card__role">React, TypeScript, дизайн-системы.</p>
                            </div>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Записаться</Link>
                            </div>
                        </div>
                        <div className="mentor-card">
                            <div className="mentor-card__top">
                                <span className="tag">JS</span>
                            </div>
                            <div className="mentor-card__image"
                                style={mentorImageStyle("https://images.unsplash.com/photo-1573496359142-b8d87734a5a2?ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80")}>
                            </div>
                            <div className="mentor-card__info">
                                <h3 className="mentor-card__title">Даниил Пахомов</h3>
                                <p className="mentor-card__role">Node.js, Nest, Next.js, БД и облака.</p>
                            </div>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Записаться</Link>
                            </div>
                        </div>
                        <div className="mentor-card">
                            <div className="mentor-card__top">
                                <span className="tag">Python</span>
                            </div>
                            <div className="mentor-card__image"
                                style={mentorImageStyle("https://images.unsplash.com/photo-1573496359142-b8d87734a5a2?ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80")}>
                            </div>
                            <div className="mentor-card__info">
                                <h3 className="mentor-card__title">Илья Громов</h3>
                                <p className="mentor-card__role">Data/ML, продакшн-пайплайны.</p>
                            </div>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Записаться</Link>
                            </div>
                        </div>
                        <div className="mentor-card">
                            <div className="mentor-card__top">
                                <span className="tag">C#</span>
                            </div>
                            <div className="mentor-card__image"
                                style={mentorImageStyle("https://images.unsplash.com/photo-1573497019940-1c28c88b4f3e?ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80")}>
                            </div>
                            <div className="mentor-card__info">
                                <h3 className="mentor-card__title">Татьяна Соколова</h3>
                                <p className="mentor-card__role">Архитектура, интеграции, безопасность.</p>
                            </div>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Записаться</Link>
                            </div>
                        </div>
                        <div className="mentor-card">
                            <div className="mentor-card__top">
                                <span className="tag">Go</span>
                            </div>
                            <div className="mentor-card__image"
                                style={mentorPlaceholderStyle}>
                                Нет фото</div>
                            <div className="mentor-card__info">
                                <h3 className="mentor-card__title">Павел Авдеев</h3>
                                <p className="mentor-card__role">gRPC/REST, resiliency, мониторинг.</p>
                            </div>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Записаться</Link>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </main>
    )
}

export default MentorsPage