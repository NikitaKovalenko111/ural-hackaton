import type React from "react"
import type { JSX } from "react/jsx-dev-runtime"
import { Link, Navigate } from "react-router-dom"
import { useSelector } from "react-redux"
import type { RootState } from "../../redux/store"

const ProfilePage: React.FC = (): JSX.Element => {
    const user = useSelector((state: RootState) => state.users.user)

    /*if (!user) {
        return <Navigate to="/login" replace />
    }*/

    return (
        <main className="container">
            <div className="profile__layout">
                <aside className="profile__card">
                    <div className="profile__header">
                        <div className="profile__avatar">👤</div>
                        <div className="profile__meta">
                            <h2>Николай Иванов</h2>
                            <p>Фронтенд-разработчик</p>
                            <span className="profile__status-badge">● Участник хаба</span>
                        </div>
                    </div>

                    <div className="info-grid">
                        <div className="info-item">
                            <span className="info-item__label">Имя</span>
                            <span className="info-item__value">Николай</span>
                        </div>
                        <div className="info-item">
                            <span className="info-item__label">Фамилия</span>
                            <span className="info-item__value">Иванов</span>
                        </div>
                        <div className="info-item">
                            <span className="info-item__label">Email</span>
                            <span className="info-item__value">n.ivanov@example.com</span>
                        </div>
                        <div className="info-item">
                            <span className="info-item__label">Telegram</span>
                            <span className="info-item__value">@niihofox</span>
                        </div>
                        <div className="info-item">
                            <span className="info-item__label">Город</span>
                            <span className="info-item__value">Москва</span>
                        </div>
                        {/* <div className="info-item">
                            <span className="info-item__label">Дата регистрации</span>
                            <span className="info-item__value">12.03.2025</span>
                        </div> */}
                    </div>

                    <h4 className="profile__section-title">О себе</h4>
                    <p className="profile__bio">
                        Изучаю React и TypeScript, участвую в open-source проектах. Ищу ментора для роста в архитектурных паттернах и
                        оптимизации рендеринга.
                    </p>

                    {/*<h4 className="profile__section-title">Интересы и стек</h4>
                    <div className="tags-container">
                        <span className="tag tag--stack">JavaScript</span>
                        <span className="tag tag--stack">React</span>
                        <span className="tag tag--stack">TypeScript</span>
                        <span className="tag tag--stack">UI/UX</span>
                        <span className="tag tag--stack">Figma</span>
                    </div>*/}

                    <div className="actions-row">
                        <Link to="/profile" className="btn btn--primary btn--sm">Редактировать профиль</Link>
                        <Link to="/profile" className="btn btn--sm btn--outline">Настройки аккаунта</Link>
                    </div>
                    <Link to="/" className="link--danger">Выйти из аккаунта</Link>
                </aside>

                <section className="notifications__section">
                    <div className="notifications__header">
                        <h2>Уведомления</h2>
                        <button className="btn btn--sm btn--outline">Отметить все как прочитанные</button>
                    </div>

                    <div className="notifications__list" id="notifList">
                        <div className="notifications__item notifications__item--unread">
                            <div className="notifications__icon">📅</div>
                            <div className="notifications__content">
                                <div className="notifications__title">Напоминание о событии</div>
                                <div className="notifications__description">Мастер-класс «Веб-дизайн» начнется через 30 минут. Не забудьте
                                    зайти в Zoom.</div>
                                <div className="notifications__time">Сегодня, 18:30</div>
                            </div>
                        </div>

                        <div className="notifications__item notifications__item--unread">
                            <div className="notifications__icon">💬</div>
                            <div className="notifications__content">
                                <div className="notifications__title">Новое сообщение от ментора</div>
                                <div className="notifications__description">Алексей Смирнов ответил на ваш вопрос по архитектуре FastAPI.
                                </div>
                                <div className="notifications__time">Сегодня, 14:15</div>
                            </div>
                        </div>

                        <div className="notifications__item">
                            <div className="notifications__icon">✅</div>
                            <div className="notifications__content">
                                <div className="notifications__title">Регистрация подтверждена</div>
                                <div className="notifications__description">Вы успешно записаны на «Практикум: CI/CD с GitHub Actions».</div>
                                <div className="notifications__time">Вчера, 20:00</div>
                            </div>
                        </div>

                        <div className="notifications__item">
                            <div className="notifications__icon">🔔</div>
                            <div className="notifications__content">
                                <div className="notifications__title">Обновление расписания</div>
                                <div className="notifications__description">Время начала лайв-кодинга React + TS перенесено на 21:00.</div>
                                <div className="notifications__time">2 дня назад</div>
                            </div>
                        </div>

                        <div className="notifications__item">
                            <div className="notifications__icon">🎓</div>
                            <div className="notifications__content">
                                <div className="notifications__title">Достижение разблокировано</div>
                                <div className="notifications__description">Вы посетили 10 мероприятий хаба! Получен бейдж «Активный
                                    участник».</div>
                                <div className="notifications__time">5 дней назад</div>
                            </div>
                        </div>
                    </div>
                </section>
            </div>
        </main>
    )
}

export default ProfilePage