import { useEffect, useState } from "react";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { fetchProcedures } from "@/lib/api";
import type { Procedure, ProcedureContentItem } from "@/types/procedure";
import { Loader2 } from "lucide-react";

export function ProceduresAccordion() {
  const [procedures, setProcedures] = useState<Procedure[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    async function loadProcedures() {
      try {
        setLoading(true);
        const data = await fetchProcedures();
        setProcedures(data);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Failed to load procedures");
      } finally {
        setLoading(false);
      }
    }

    loadProcedures();
  }, []);

    const parseContent = (content: string): ProcedureContentItem[] => {
        try {
            const parsed = JSON.parse(content);
            if (!Array.isArray(parsed)) return [];
            return parsed.map((item) => {
                if (typeof item === "string") {
                    return { value: item };
                }
                if (typeof item === "object" && item !== null && "value" in item) {
                    return item as ProcedureContentItem;
                }
                return null;
            }).filter(Boolean) as ProcedureContentItem[];
        } catch {
            return content ? [{ value: content }] : [];
        }
    };


    const renderContent = (content: string) => {
        const items = parseContent(content);

        if (items.length === 0) {
            return (
                <p className="text-muted-foreground italic">
                    Нет дополнительной информации
                </p>
            );
        }

        return (
            <ul className="space-y-3">
                {items.map((item, index) => (
                    <li
                        key={index}
                        className="flex items-start gap-3 rounded-lg border bg-background px-4 py-3 shadow-sm hover:shadow transition"
                    >
                        <span className="mt-1 text-primary">✔</span>

                        <span className="flex-1 leading-relaxed">
            {item.url ? (
                <a
                    href={item.url}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-primary font-medium hover:underline"
                >
                    {item.value}
                </a>
            ) : (
                item.value
            )}
          </span>
                    </li>
                ))}
            </ul>
        );
    };


  const defaultValues = procedures
    .filter((p) => p.is_expanded)
    .map((p) => `item-${p.id}`);

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-[400px]">
        <Loader2 className="h-8 w-8 animate-spin text-primary" />
      </div>
    );
  }

  if (error) {
    return (
      <Card className="border-destructive">
        <CardContent className="pt-6">
          <p className="text-destructive">Ошибка: {error}</p>
        </CardContent>
      </Card>
    );
  }

  return (
    <div className="w-full max-w-5xl mx-auto p-4 sm:p-6">
      <Card className="shadow-xl border-2">
        <CardHeader className="bg-gradient-to-r from-primary/5 to-primary/10 border-b">
          <CardTitle className="text-xl sm:text-2xl font-bold text-center tracking-tight">
            ПРОЦЕДУРЫ РОЗЫСКА ПОСЫЛОК, А ТАКЖЕ ОФОРМЛЕНИЯ ЗАЯВЛЕНИЙ О ПОВРЕЖДЕНИИ ИЛИ УТРАТЕ
          </CardTitle>
        </CardHeader>
        <CardContent className="pt-6">
          <Accordion
            type="multiple"
            defaultValue={defaultValues}
            className="w-full space-y-3"
          >
            {procedures.map((procedure) => (
              <AccordionItem
                key={procedure.id}
                value={`item-${procedure.id}`}
                className="border-2 rounded-xl px-5 bg-card hover:bg-accent/30 transition-all duration-200 shadow-sm hover:shadow-md"
              >
                <AccordionTrigger className="text-left font-semibold hover:no-underline py-5 text-base sm:text-lg">
                  <span className="flex-1 text-left pr-4">{procedure.title}</span>
                </AccordionTrigger>
                <AccordionContent className="text-sm sm:text-base pb-5 pt-2 text-muted-foreground">
                  <div className="pl-2">
                    {renderContent(procedure.content)}
                  </div>
                </AccordionContent>
              </AccordionItem>
            ))}
          </Accordion>
        </CardContent>
      </Card>
    </div>
  );
}

