-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS procedures (
                                          id SERIAL PRIMARY KEY,
                                          title VARCHAR(500) NOT NULL,
                                          type VARCHAR(100) NOT NULL,
                                          content TEXT NOT NULL DEFAULT '[]',
                                          sort_order INTEGER NOT NULL DEFAULT 0,
                                          is_expanded BOOLEAN NOT NULL DEFAULT false,
                                          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                          updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_procedures_type ON procedures(type);
CREATE INDEX IF NOT EXISTS idx_procedures_sort_order ON procedures(sort_order);

INSERT INTO procedures (title, type, content, sort_order, is_expanded) VALUES
                                                                           (
                                                                               'Необходимые документы для розыска посылок без статуса доставки',
                                                                               'search_without_status',
                                                                               '[
                                                                                 "коммерческий инвойс (оформленный на клиента)",
                                                                                 "накладная",
                                                                                 "либо подробное описание содержимого отправления"
                                                                               ]',
                                                                               1,
                                                                               true
                                                                           ),
                                                                           (
                                                                               'Необходимые документы в случае утраты или повреждения посылки',
                                                                               'loss_or_damage_docs',
                                                                               '[
                                                                                 "заявление на розыск",
                                                                                 "копия паспорта получателя",
                                                                                 "подтверждение стоимости вложения"
                                                                               ]',
                                                                               2,
                                                                               false
                                                                           ),
                                                                           (
                                                                               'Дополнительные документы в случае повреждения',
                                                                               'damage_additional_docs',
                                                                               '[
                                                                                 "фотографии повреждений",
                                                                                 "акт осмотра",
                                                                                 "упаковка с маркировкой"
                                                                               ]',
                                                                               3,
                                                                               false
                                                                           ),
                                                                           (
                                                                               'Порядок действий в случае утраты посылки',
                                                                               'loss_procedure',
                                                                               '[
                                                                                 "обратиться в службу поддержки",
                                                                                 "подать заявление на розыск",
                                                                                 "ожидать ответа до 30 дней"
                                                                               ]',
                                                                               4,
                                                                               false
                                                                           ),
                                                                           (
                                                                               'Порядок действий в случае повреждения посылки',
                                                                               'damage_procedure',
                                                                               '[
                                                                                 "зафиксировать повреждение при получении",
                                                                                 "сделать фото",
                                                                                 "обратиться в службу поддержки"
                                                                               ]',
                                                                               5,
                                                                               false
                                                                           ),
                                                                           (
                                                                               'Важная информация для получателя',
                                                                               'recipient_info',
                                                                               '[
                                                                                 "проверяйте посылку при получении",
                                                                                 "сохраняйте упаковку до проверки содержимого"
                                                                               ]',
                                                                               6,
                                                                               false
                                                                           );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS procedures;
-- +goose StatementEnd
