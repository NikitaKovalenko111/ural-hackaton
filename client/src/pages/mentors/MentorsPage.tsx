import { isAxiosError } from "axios"
import { useEffect, useState } from "react"
import type { CSSProperties, JSX } from "react"
import RegistrationModal from "../../components/registrationModal/RegistrationModal"
import { mentorsApi } from "../../api/api"
import type { IMentor } from "../../types"

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

const mentorImages = [
    "https://images.unsplash.com/photo-1531891437562-4301cf35b7e4?ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80",
    "https://images.unsplash.com/photo-1544716278-ca5e3f4abd8c?ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80",
    "https://images.unsplash.com/photo-1560250097-0b93528c311a?ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80",
    "https://images.unsplash.com/photo-1573496359142-b8d87734a5a2?ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80",
]

const MentorsPage = (): JSX.Element => {
    const [isRegistrationOpen, setIsRegistrationOpen] = useState<boolean>(false)
    const [selectedMentorName, setSelectedMentorName] = useState<string>("Выбранный наставник")
    const [selectedMentorId, setSelectedMentorId] = useState<number | undefined>(undefined)
    const [mentors, setMentors] = useState<IMentor[]>([])
    const [isLoading, setIsLoading] = useState<boolean>(true)
    const [errorMessage, setErrorMessage] = useState<string>("")

    useEffect(() => {
        const loadMentors = async (): Promise<void> => {
            setIsLoading(true)
            setErrorMessage("")

            try {
                const data = await mentorsApi.getMentors()
                setMentors(data)
            } catch (error) {
                setMentors([])
                if (isAxiosError(error) && typeof error.response?.data?.message === "string") {
                    setErrorMessage(error.response.data.message)
                } else {
                    setErrorMessage("Не удалось загрузить наставников")
                }
            } finally {
                setIsLoading(false)
            }
        }

        void loadMentors()
    }, [])

    const handleMentorRegistrationClick = (mentorName: string, mentorId?: number): void => {
        setSelectedMentorName(mentorName)
        setSelectedMentorId(mentorId)
        setIsRegistrationOpen(true)
    }

    return (
        <main>
            <section className="section section--mentors">
                <div className="container">
                    <div className="section__header">
                        <div>
                            <h2 className="section__title">Наставники</h2>
                            <p className="section__subtitle">Список загружается с сервера. Выберите ментора и оставьте заявку.</p>
                        </div>
                    </div>

                    <div className="stats__container" style={statsContainerStyle}>
                        <div className="stats__box" style={statsBoxStyle}>
                            <div className="stats__label">Менторов доступно</div>
                            <div className="stats__value">{mentors.length}</div>
                            <div className="stats__unit">👥</div>
                        </div>
                    </div>

                    <div className="mentors-grid">
                        {mentors.map((mentor, index) => (
                            <div className="mentor-card" key={`${mentor.mentorId ?? mentor.id}-${mentor.email}`}>
                                <div className="mentor-card__top">
                                    <span className="tag">Ментор</span>
                                </div>
                                {mentorImages[index % mentorImages.length] ? (
                                    <div className="mentor-card__image" style={mentorImageStyle(mentorImages[index % mentorImages.length])} />
                                ) : (
                                    <div className="mentor-card__image" style={mentorPlaceholderStyle}>Нет фото</div>
                                )}
                                <div className="mentor-card__info">
                                    <h3 className="mentor-card__title">{mentor.fullname}</h3>
                                    <p className="mentor-card__role">Telegram: {mentor.telegram || "не указан"}</p>
                                    <p className="mentor-card__role">Email: {mentor.email}</p>
                                </div>
                                <div className="mentor-card__actions">
                                    <button type="button" className="btn btn--primary btn--sm" onClick={() => handleMentorRegistrationClick(mentor.fullname, mentor.mentorId)}>
                                        Записаться
                                    </button>
                                </div>
                            </div>
                        ))}
                    </div>

                    {isLoading ? <p className="hubs-search__empty">Загрузка наставников...</p> : null}
                    {errorMessage ? <p className="hubs-search__empty">{errorMessage}</p> : null}
                    {!isLoading && !errorMessage && mentors.length === 0 ? (
                        <p className="hubs-search__empty">Наставники не найдены</p>
                    ) : null}
                </div>
            </section>

            <RegistrationModal
                isOpen={isRegistrationOpen}
                onClose={() => {
                    setIsRegistrationOpen(false)
                    setSelectedMentorId(undefined)
                }}
                eventTitle={selectedMentorName}
                registrationType="mentor"
                mentorId={selectedMentorId}
            />
        </main>
    )
}

export default MentorsPage
