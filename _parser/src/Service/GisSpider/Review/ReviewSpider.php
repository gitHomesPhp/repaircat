<?php

declare(strict_types=1);

namespace App\Service\GisSpider\Review;

use App\Entity\Sc;
use App\Entity\Url;
use DateTimeImmutable;
use Doctrine\ORM\EntityManagerInterface;
use GuzzleHttp\Client;
use stdClass;

class ReviewSpider
{
    public function __construct(
        private Client $client,
        private EntityManagerInterface $em,
    ) {}

    public function handle(?int $projectScId = null)
    {
        if ($projectScId) {
            $scRepo = $this->em->getRepository(Sc::class);
            $sc = $scRepo->findOneBy(['repaircat_id' => $projectScId]);

            //TODO не факт что в базе в будущем будет инт
            $_2gisId = (int)$sc->getGisId();

            $link = new Link();
            $this->dig($link->setItemId($_2gisId)->buildUrl());
        }

    }

    public function dig(Link $link): void
    {
        $response = $this->client->get($link->getUrl());

        $url = new Url();
        $url->setHost(Link::HOST);
        $url->setType(Url::REVIEW_TYPE);
        $url->setPath($link->getPath());
        $url->setLastRequest(new DateTimeImmutable());
        $data = $response->getBody()->getContents();

        if ($response->getStatusCode() === 200) {
            $url->setIsResponseOK(true);
            $url->setData($data);
        }

        $this->em->persist($url);
        $this->em->flush();

        $this->getScReviewNextLink(json_decode($data));
    }

    private function getScReviewNextLink(stdClass $dataObject): void
    {
        sleep(mt_rand(1, 5));
        if (isset($dataObject->meta->next_link)) {
            $link = new Link();
            $link->setUrl($dataObject->meta->next_link);

            $this->dig($link);
        }
    }
}